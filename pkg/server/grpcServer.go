package server

import (
	"context"
	"fmt"
	"sixletters/simple-db/pkg/storage"
	"sixletters/simple-db/pkg/tree"
	"sixletters/simple-db/pkg/stubs"
)

type grpcServer struct{
	storageEngine storage.StorageEngine
}

func NewGrpcServer() *grpcServer {
	newStorageConfig := storage.NewConfigWithOpts(
		storage.WithTreeType(tree.BTree),
		storage.WithFilePath("testFile.txt"),
	)
	storageEngine, err := storage.NewSingletonEngine(newStorageConfig)
	if err != nil {
		panic(err)
	}
	return &grpcServer{
		storageEngine: storageEngine,
	}
}

func (s *grpcServer) Get(ctx context.Context, req *stubs.GetRequest) (*stubs.GetResponse, error) {
	res, err := s.storageEngine.Get(context.Background(), req.Key)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return &stubs.GetResponse{Value: res}, nil
}

func (s *grpcServer) Put(ctx context.Context, req *stubs.PutRequest) (*stubs.PutResponse, error) {
	err := s.storageEngine.Put(context.Background(), req.Key, req.Value)
	if err != nil {
		fmt.Print(err.Error())
		return &stubs.PutResponse{Success: false}, err
	}
	return &stubs.PutResponse{Success: true}, nil
}
