package main

import (
	"flag"
	"sixletters/simple-db/pkg/config"
	httpserver "sixletters/simple-db/pkg/http_server"
	"sixletters/simple-db/pkg/storage"
	"sixletters/simple-db/pkg/tree"
)

var host string
var port string
var treeType string
var dataDir string

func init() {
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.StringVar(&host, "host", "", "host to listen on")
	// other trees not supported yet
	// flag.StringVar(&treeType, "tree-type", "", "The type of tree used for the storage engine")
	flag.StringVar(&dataDir, "data-dir", "/Users/simpledb/data", "The path to the data directory")
}

// Todo: use cobra command to start in the future
func main() {
	flag.Parse()

	engineConfig := config.StorageEngineConfig{
		DataDir:  dataDir,
		TreeType: tree.BTree,
	}
	if err := storage.InitStorageEngine(&engineConfig); err != nil {
		panic(err)
	}

	serverConfig := config.HttpServerConfigs{
		Host: host,
		Port: port,
	}
	httpServer := httpserver.NewHttpServer(&serverConfig)
	if err := httpServer.Run(); err != nil {
		panic(err)
	}

}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"sixletters/simple-db/pkg/config"
// 	"sixletters/simple-db/pkg/storage"
// 	"sixletters/simple-db/pkg/tree"
// )

// func main() {
// 	fmt.Println("HELLO WORLD")
// 	newStorageConfig := config.NewStorageConfigWithOpts(
// 		config.WithTreeType(tree.BTree),
// 		config.WithDataDir(""),
// 	)
// 	storageEngine, err := storage.NewSingletonEngine(newStorageConfig)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = storageEngine.Put(context.Background(), "Key1", "DION")
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	err = storageEngine.Put(context.Background(), "Key2", "HARRIS")
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	val, err := storageEngine.Get(context.Background(), "Key1")
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	fmt.Print(val)
// }
