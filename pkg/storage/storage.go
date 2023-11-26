package storage

import (
	"context"
	"sixletters/simple-db/pkg/tree"
)

type EngineOpt func(StorageEngine)
type EngineConfigOpt func(*StorageEngineConfig)

// The storage engine interface is a set of traits that any implementation of this storage engine should possess
type StorageEngine interface {
	Put(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
}

// A structure that holds the configs for storage engine intializations, will be used in both the singleton and distributed verion
type StorageEngineConfig struct {
	Filepath string
	TreeType tree.TreeType
}

func WithFilePath(filepath string) EngineConfigOpt {
	return func(config *StorageEngineConfig) {
		config.Filepath = filepath
	}
}

func WithTreeType(treeType tree.TreeType) EngineConfigOpt {
	return func(config *StorageEngineConfig) {
		config.TreeType = treeType
	}
}

func NewConfigWithOpts(opts ...EngineConfigOpt) *StorageEngineConfig {
	config := &StorageEngineConfig{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}
