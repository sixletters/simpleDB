package main

import (
	"fmt"
	"log"
	"os"
	"sixletters/simple-db/pkg/btree"
)

func main() {
	fmt.Println("HELLO WORLD")
	f, err := os.Create("testFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mytree := btree.NewBTree(f)
	mytree.Insert("b", "dion")
	mytree.Insert("c", "dion")
	mytree.Insert("a", "dion")
	mytree.Insert("d", "dion")
	mytree.Insert("e", "dion")
	mytree.Insert("g", "dion")
	fmt.Println(mytree.Search("a"))
	// mytree.Insert("f", "dion")
	// mytree.Insert("h", "dion")
	// mytree.PrintTree()
}
