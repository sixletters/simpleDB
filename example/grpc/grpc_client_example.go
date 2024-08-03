package main

import (
	"context"
	"fmt"
	"sixletters/simple-db/pkg/client"
)

func main() {
	newGrpcClient, err := client.NewGrpcClient("127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	val, err := newGrpcClient.Get(context.Background(), "test")
	fmt.Print(val)
}
