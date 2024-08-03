package client

import (
	"context"
	proto "sixletters/simple-db/pkg/proto/interface"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient proto.SimpleDBClient
}

func NewGrpcClient(serverAddr string) (SimpleDBClient, error) {
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := proto.NewSimpleDBClient(conn)
	return &GrpcClient{
		conn:       conn,
		grpcClient: client,
	}, nil
}

func (c *GrpcClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *GrpcClient) Get(ctx context.Context, Key string) (string, error) {
	val, err := c.grpcClient.Get(ctx, &proto.Key{Key: Key})
	return val.Value, err
}

func (c *GrpcClient) Put(ctx context.Context, Key string, Value string) (string, error) {
	val, err := c.grpcClient.Put(ctx, &proto.KeyValue{Key: Key, Value: Value})
	return val.Value, err
}
