package main

import (
	"fmt"
	"log"
	"os"
	"sixletters/simple-db/pkg/btree"
	"sixletters/simple-db/pkg/token"
)

func main() {
	fmt.Println("HELLO WORLD")
	f, err := os.OpenFile("testFile.txt", os.O_RDWR, os.ModeAppend)
	// f, err := os.Create("testFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	info, err := f.Stat()
	fmt.Println(info.Size())
	mytree, _ := btree.NewBTree(f)
	// mytree.Insert("b", "dion")
	// mytree.Insert("c", "dion")
	// mytree.Insert("a", "dion")
	// mytree.Insert("d", "dion")
	// mytree.Insert("e", "Harris")
	// mytree.Insert("g", "Larris")
	// mytree.Insert("g", "dion")
	// mytree.Insert("g", "dion")
	// fmt.Println(info.Size())
	fmt.Println(mytree.Search("e"))
	// mytree.Insert("f", "dion")
	// mytree.Insert("h", "dion")
	// mytree.PrintTree()
	mytoken, _ := token.GenerateToken("HARRIS")
	fmt.Print(mytoken)
	if err := token.Authenticate(mytoken, "HARRIS"); err != nil {
		fmt.Print(err.Error())
		panic("TOKEN IS WRONG")
	}
	fmt.Printf("SUCCESS")
}
