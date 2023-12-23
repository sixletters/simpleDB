package storage

import (
	"context"
	"fmt"
	"os"
	"sixletters/simple-db/pkg/config"
	"sixletters/simple-db/pkg/tree"
	"sync"
)

// The singleton Engine is a simple implementation of the StorageEngine Interface. It is a single standalone
// storage engine that uses a tree as the underlying storage data structure. The tree type can be passed in into
// the storageEngineConfigs
type singletonEngine struct {
	File     *os.File
	Tree     tree.Tree
	TreeLock *sync.RWMutex
	Config   *config.StorageEngineConfig
}

// Constructor for the singleton Engine
func NewSingletonEngine(config *config.StorageEngineConfig, opts ...EngineOpt) (StorageEngine, error) {
	// logic should and can be abstracted out
	var engineFile *os.File
	// only retrive directory from user, the naming of the file will be fixed
	dataFilePath := getDataFilePath(config.DataDir)
	_, err := os.Stat(dataFilePath)
	if err != nil {
		// create file it does not exist yet
		engineFile, err = os.Create(dataFilePath)
		if err != nil {
			return nil, err
		}
	} else {
		engineFile, err = os.OpenFile(dataFilePath, os.O_RDWR, os.ModeAppend)
		if err != nil {
			return nil, err
		}
	}
	engine := &singletonEngine{
		Config:   config,
		TreeLock: &sync.RWMutex{},
		File:     engineFile,
	}
	for _, opt := range opts {
		opt(engine)
	}
	switch config.TreeType {
	case tree.BTree:
		engineTree, err := tree.NewTree(tree.BTree, engineFile)
		if err != nil {
			return nil, err
		}
		engine.Tree = engineTree
	default:
		panic("Treetype not supported")
	}
	return engine, nil
}

// Simple put for the engine
func (se *singletonEngine) Put(ctx context.Context, key string, value string) error {
	se.TreeLock.Lock()
	defer se.TreeLock.Unlock()

	fmt.Printf("Putting key: %s", key)
	return se.Tree.Insert(key, value)
}

// Simple get for the engine
func (se *singletonEngine) Get(ctx context.Context, key string) (string, error) {
	se.TreeLock.RLock()
	defer se.TreeLock.RUnlock()
	val, found := se.Tree.Search(key)
	if !found {
		return "", fmt.Errorf("no value was found for key: %s", key)
	}
	return val, nil
}
