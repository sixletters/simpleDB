package storage

import "context"

type StorageEngine interface {
	Put(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type storageEngine struct {
}

func NewStorageEngine() StorageEngine {
	return &storageEngine{}
}

func (s *storageEngine) Put(ctx context.Context, key string, value string) error {
	return nil
}

func (s *storageEngine) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}
