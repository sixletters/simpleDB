package tree

import (
	"os"
	"sixletters/simple-db/pkg/tree/block"
	"sixletters/simple-db/pkg/tree/btree"
)

type TreeType int
type TreeOptions func(Tree)

// This is an interface that is used to implement any form of storage tree
// data structure
type Tree interface {
	GetBlockManager() block.BlockManager
	GetBlockFile() *os.File
	PrintTree()
	Search(key string) (string, bool)
	Insert(key string, value string) error
}

// Currently only Btrees are supported
const (
	BTree TreeType = iota
	LstmTree
	BpTree
)

// NewTree returns a new instance of a tree given a tree type, file pointer and options
func NewTree(treeType TreeType, file *os.File, options ...TreeOptions) (Tree, error) {
	var tree Tree
	var err error
	switch treeType {
	case BTree:
		tree, err = btree.NewBTree(file)
		if err != nil {
			return nil, err
		}
	default:
		panic("Tree type no supported")
	}
	for _, opt := range options {
		opt(tree)
	}
	return tree, nil
}
