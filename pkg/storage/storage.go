package storage

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sixletters/simple-db/pkg/config"
	"sixletters/simple-db/pkg/consts"
	"sixletters/simple-db/pkg/tree"
)

type EngineOpt func(StorageEngine)

// the storage engine exists as a singleton in this file scope
// todo: evaluate if this is a good idea
var storageEngine StorageEngine

// The storage engine interface is a set of traits that any implementation of this storage engine should possess
type StorageEngine interface {
	Put(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
}

// This inits the appropriate singleton engine with the appropriate tree
func InitStorageEngine(config *config.StorageEngineConfig) error {
	if storageEngine != nil {
		return fmt.Errorf("engine has already been inited")
	}
	var (
		engine StorageEngine
		err    error
	)
	// Create the data directory for the storage engines if it does not exist
	if _, err := os.Stat(config.DataDir); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(config.DataDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	// Create the type of engine
	switch config.TreeType {
	case tree.BTree:
		engine, err = NewSingletonEngine(config)
		if err != nil {
			return err
		}
		storageEngine = engine
	default:
		return fmt.Errorf("unsupported tree type")
	}
	return nil
}

func Get(ctx context.Context, key string) (string, error) {
	return storageEngine.Get(ctx, key)
}

func Put(ctx context.Context, key string, value string) error {
	return storageEngine.Put(ctx, key, value)
}

// todo: validate in the future
func getDataFilePath(dataDir string) string {
	return fmt.Sprintf("%s/%s", dataDir, consts.DataFileName)
}
