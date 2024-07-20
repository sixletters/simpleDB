package client

import "context"

type SimpleDBClient interface {
	Get(ctx context.Context, Key string) (string, error)
	Put(ctx context.Context, Key string, Value string) (string, error)
	Close() error
}
