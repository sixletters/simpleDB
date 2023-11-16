package server

import (
	"context"
)

type KeyValueServer struct{}

// mustEmbedUnimplementedKeyValueServiceServer implements KeyValueServiceServer.
func (*KeyValueServer) mustEmbedUnimplementedKeyValueServiceServer() {
	panic("unimplemented")
}

func (s *KeyValueServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	// Implement your logic for the Get method here
	// For simplicity, this example always returns a hard-coded value "value for key"
	return &GetResponse{Value: "value for key"}, nil
}

func (s *KeyValueServer) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	// Implement your logic for the Put method here
	// For simplicity, this example always returns success as true
	return &PutResponse{Success: true}, nil
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
