package httpserver

import (
	"bytes"
	"fmt"
	"io"
	"sixletters/simple-db/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// Response basic response for all response
type Response struct {
	ErrCode int         `json:"error_code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Success return success status
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, &Response{0, "success", data})
}

// Failed return failed status
func Failed(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, &Response{code, message, data})
}

func FailedWithInternalError(c *gin.Context, err error) {
	Failed(c, 500, err.Error(), nil)
}

func FailedWithParams(c *gin.Context, err error) {
	Failed(c, 400, err.Error(), nil)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logs the handler requests
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			glog.Error(err.Error())
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		glog.Infof("Request: %s %s, body: %s\n", c.Request.Method, c.Request.URL.String(), reqBody)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			glog.Infof("Response of %s %s: %s\n", c.Request.Method, c.Request.URL.String(), blw.body.String())
		}
	}
}

func validateGetQuery(query *model.GetQueryRequest) error {
	if query.Key == "" {
		return fmt.Errorf("key cannot be an empty string")
	}
	return nil
}

func validatePutQuery(query *model.PutRequest) error {
	if query.Key == "" {
		return fmt.Errorf("key cannot be an empty string")
	}
	return nil
}
