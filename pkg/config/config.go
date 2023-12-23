package config

import "sixletters/simple-db/pkg/tree"

type EngineConfigOpt func(*StorageEngineConfig)

type HttpServerConfigs struct {
	Host string
	Port string
}

// A structure that holds the configs for storage engine intializations, will be used in both the singleton and distributed verion
type StorageEngineConfig struct {
	DataDir  string
	TreeType tree.TreeType
}

func WithDataDir(dataDir string) EngineConfigOpt {
	return func(config *StorageEngineConfig) {
		config.DataDir = dataDir
	}
}

func WithTreeType(treeType tree.TreeType) EngineConfigOpt {
	return func(config *StorageEngineConfig) {
		config.TreeType = treeType
	}
}

func NewStorageConfigWithOpts(opts ...EngineConfigOpt) *StorageEngineConfig {
	config := &StorageEngineConfig{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}
