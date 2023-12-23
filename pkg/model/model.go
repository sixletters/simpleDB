package model

type GetQueryRequest struct {
	Key string `json:"key"`
}

type GetResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PutRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PutResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
