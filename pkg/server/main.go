package server

import (
	"context"
	"errors"
	"log"
	"sync"
	"fmt"
	"sixletters/simple-db/pkg/storage"
	"sixletters/simple-db/pkg/tree"
)

type KeyValueServer struct{
	storageEngine *storage.StorageEngine
}

type notFoundError struct {
	Key string
}
	
// mustEmbedUnimplementedKeyValueServiceServer implements KeyValueServiceServer.
func (*KeyValueServer) mustEmbedUnimplementedKeyValueServiceServer() {
	panic("unimplemented")
}

func NewKeyValueServer() *KeyValueServer {
	fmt.Println("HELLO WORLD")
	newStorageConfig := storage.NewConfigWithOpts(
		storage.WithTreeType(tree.BTree),
		storage.WithFilePath("testFile.txt"),
	)
	storageEngine, err := storage.NewSingletonEngine(newStorageConfig)
	if err != nil {
		panic(err)
	}
	return &KeyValueServer{
		storageEngine: &storageEngine,
	}
}

func (s *KeyValueServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	// Implement your logic for the Get method here
	// For simplicity, this example always returns a hard-coded value "value for key"
	// s.mu.Lock()
	// defer s.mu.Unlock()

	// value, exists := s.store[req.Key]
	// if !exists {
	// 	return nil, NotFoundError(req.Key)
	// }


	// return &GetResponse{Value: value}, nil
	res, err := (*s.storageEngine).Get(context.Background(), req.Key)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	return &GetResponse{Value: res}, nil


}

func (s *KeyValueServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	// Implement your logic for the Put method here
	// For simplicity, this example always returns success as true
	// log.Println("Entering Put method")
	// s.mu.Lock()
	// defer s.mu.Unlock()

	// // Check for nil pointers or other potential issues
	// if req == nil {
	// 	log.Println("Received nil PutRequest")
	// 	return nil, errors.New("nil PutRequest")
	// }

	// s.store[req.Key] = req.Value


	// return &PutResponse{Success: true}, nil
	err := (*s.storageEngine).Put(context.Background(), req.Key, req.Value)
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

// func main() {
// 	lis, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	RegisterKeyValueServiceServer(grpcServer, &keyValueServer{})

// 	log.Println("Server is listening on :8080")
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to serve: %v", err)
// 	}
// }
