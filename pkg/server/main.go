package server

import (
	"context"
	"errors"
	"log"
	"sync"
)

type KeyValueServer struct{
	mu sync.Mutex
	store map[string]string
}

type notFoundError struct {
	Key string
}
	
// mustEmbedUnimplementedKeyValueServiceServer implements KeyValueServiceServer.
func (*KeyValueServer) mustEmbedUnimplementedKeyValueServiceServer() {
	panic("unimplemented")
}

func NewKeyValueServer() *KeyValueServer {
	return &KeyValueServer{
		store: make(map[string]string),
	}
}

func (s *KeyValueServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	// Implement your logic for the Get method here
	// For simplicity, this example always returns a hard-coded value "value for key"
	s.mu.Lock()
	defer s.mu.Unlock()

	value, exists := s.store[req.Key]
	if !exists {
		return nil, NotFoundError(req.Key)
	}


	return &GetResponse{Value: value}, nil
}

func (s *KeyValueServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	// Implement your logic for the Put method here
	// For simplicity, this example always returns success as true
	log.Println("Entering Put method")
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check for nil pointers or other potential issues
	if req == nil {
		log.Println("Received nil PutRequest")
		return nil, errors.New("nil PutRequest")
	}

	s.store[req.Key] = req.Value


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
