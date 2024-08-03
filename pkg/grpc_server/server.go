package grpcserver

import (
	"context"
	"fmt"
	"net"
	proto "sixletters/simple-db/pkg/proto/interface"
	"sixletters/simple-db/pkg/storage"

	"google.golang.org/grpc"
)

type simpleDBGrpcServer struct {
	host string
	port string
	proto.UnimplementedSimpleDBServer
}

func NewGrpcServer(host string, port string) *simpleDBGrpcServer {
	return &simpleDBGrpcServer{
		host: host,
		port: port,
	}
}

func (s *simpleDBGrpcServer) Get(ctx context.Context, payload *proto.Key) (*proto.Value, error) {
	val, err := storage.Get(ctx, payload.Key)
	if err != nil {
		return nil, err
	}
	return &proto.Value{Value: val}, nil
}

func (s *simpleDBGrpcServer) Put(ctx context.Context, payload *proto.KeyValue) (*proto.KeyValue, error) {
	if err := storage.Put(ctx, payload.Key, payload.Value); err != nil {
		return nil, err
	}
	return payload, nil
}

func (s *simpleDBGrpcServer) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterSimpleDBServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
