package httpserver

import (
	"fmt"
	"sixletters/simple-db/pkg/config"
	"sixletters/simple-db/pkg/model"
	"sixletters/simple-db/pkg/storage"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	host    string
	port    string
	handler *gin.Engine
}

// constructor for the http server
func NewHttpServer(cfg *config.HttpServerConfigs) *HttpServer {
	server := &HttpServer{
		host: cfg.Host,
		port: cfg.Port,
	}
	if err := server.setupHandlers(); err != nil {
		panic(err)
	}
	return server
}

func (s *HttpServer) setupHandlers() error {
	handler := gin.Default()
	handler.Use(gin.Recovery())
	handler.Use(RequestLogger())

	// HTTP handlers
	apiGroup := handler.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			v1Group.GET("/get", s.Get)
			v1Group.POST("/put", s.Put)
		}
	}
	s.handler = handler
	return nil
}

func (s *HttpServer) Run() error {
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	return s.handler.Run(addr)
}

func (s *HttpServer) Put(c *gin.Context) {
	req := &model.PutRequest{}
	if err := c.Bind(req); err != nil {
		FailedWithParams(c, err)
		return
	}
	if err := validatePutQuery(req); err != nil {
		FailedWithParams(c, err)
		return
	}
	if err := storage.Put(c, req.Key, req.Value); err != nil {
		FailedWithInternalError(c, err)
		return
	}
	Success(c, model.PutResponse{Key: req.Key, Value: req.Value})
}

func (s *HttpServer) Get(c *gin.Context) {
	var err error
	key := c.Query("key")
	value := ""
	query := &model.GetQueryRequest{Key: key}
	if err := c.Bind(query); err != nil {
		FailedWithParams(c, err)
		return
	}
	if err := validateGetQuery(query); err != nil {
		FailedWithParams(c, err)
		return
	}
	if value, err = storage.Get(c, query.Key); err != nil {
		FailedWithInternalError(c, err)
		return
	}
	Success(c, model.GetResponse{Key: query.Key, Value: value})
}
