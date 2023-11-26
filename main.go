package main

import (
	"context"
	"fmt"
	"sixletters/simple-db/pkg/storage"
	"sixletters/simple-db/pkg/tree"
)

func main() {
	fmt.Println("HELLO WORLD")
	newStorageConfig := storage.NewConfigWithOpts(
		storage.WithTreeType(tree.BTree),
		storage.WithFilePath("testFile.txt"),
	)
	storageEngine, err := storage.NewSingletonEngine(newStorageConfig)
	if err != nil {
		panic(err)
	}

	err = storageEngine.Put(context.Background(), "Key1", "DION")
	if err != nil {
		fmt.Print(err.Error())
	}
	err = storageEngine.Put(context.Background(), "Key2", "HARRIS")
	if err != nil {
		fmt.Print(err.Error())
	}
	val, err := storageEngine.Get(context.Background(), "Key1")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(val)
}
