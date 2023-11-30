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

// "fmt"
// "sixletters/simple-db/pkg/simpleDickbig"
// "google.golang.org/grpc"
// func main() {
//     lis, err := net.Listen("tcp", ":8080")
//     if err != nil {
//         log.Fatalf("Failed to listen: %v", err)
//     }

//     grpcServer := grpc.NewServer()
//     simpleDickbig.RegisterKeyValueServiceServer(grpcServer, &simpleDickbig.keyValueServer{})

//     log.Println("Server is listening on :8080")
//     if err := grpcServer.Serve(lis); err != nil {
//         log.Fatalf("Failed to serve: %v", err)
//     }
// }
