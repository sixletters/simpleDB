package server

import (
	"context"
	"fmt"
	"sixletters/simple-db/pkg/storage"
	"sixletters/simple-db/pkg/tree"
)

type KeyValueServer struct{
	storageEngine storage.StorageEngine
}

type notFoundError struct {
	Key string
}
	
// mustEmbedUnimplementedKeyValueServiceServer implements KeyValueServiceServer.
func (*KeyValueServer) mustEmbedUnimplementedKeyValueServiceServer() {
	panic("unimplemented")
}

func NewKeyValueServer() *KeyValueServer {
	newStorageConfig := storage.NewConfigWithOpts(
		storage.WithTreeType(tree.BTree),
		storage.WithFilePath("testFile.txt"),
	)
	storageEngine, err := storage.NewSingletonEngine(newStorageConfig)
	if err != nil {
		panic(err)
	}
	return &KeyValueServer{
		storageEngine: storageEngine,
	}
}

func (s *KeyValueServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	res, err := s.storageEngine.Get(context.Background(), req.Key)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return &GetResponse{Value: res}, nil


}

func (s *KeyValueServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	err := s.storageEngine.Put(context.Background(), req.Key, req.Value)
	if err != nil {
		fmt.Print(err.Error())
		return &PutResponse{Success: false}, err
	}
	return &PutResponse{Success: true}, nil
}

// NotFoundError creates an error indicating that the requested key was not found.
func NotFoundError(key string) error {
	return &notFoundError{Key: key}
}

func (e *notFoundError) Error() string {
	return "key not found: " + e.Key
}
