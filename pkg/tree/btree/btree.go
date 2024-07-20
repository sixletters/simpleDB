package btree

import (
	"fmt"
	"os"
	"sixletters/simple-db/pkg/consts"
	blockUtil "sixletters/simple-db/pkg/tree/block"
)

// consts "sixletters/kv-store/pkg/consts"

type BTree struct {
	Root             *blockUtil.Pnode
	minItems         int
	treeBlockManager blockUtil.BlockManager
}

func NewBTree(file *os.File) (*BTree, error) {
	bManager := blockUtil.NewNaiveBlockManager(file)
	btree := &BTree{}
	if bManager.RootBlockExists() {
		rootBlock, err := bManager.GetRootBlock()
		if err != nil {
			fmt.Println(err.Error())
			panic("unable to get root block, file may be corrupted")
		}
		btree.Root = blockUtil.NewPnode(bManager).FromBlock(rootBlock)
	} else {
		btree.Root = blockUtil.NewPnode(bManager).WithID(0)
	}
	btree.treeBlockManager = bManager
	btree.minItems = consts.DefaultMinimumItems
	return btree, btree.Root.Commit()
}

func (bt *BTree) GetBlockManager() blockUtil.BlockManager {
	return bt.treeBlockManager
}

func (bt *BTree) GetBlockFile() *os.File {
	return bt.treeBlockManager.GetBlockFile()
}

func (bt *BTree) PrintTree() {
	if bt.Root != nil {
		bt.Root.PrintTree()
	}
}

func (bt *BTree) Search(key string) (string, bool) {
	if bt.Root == nil {
		return "", false
	}
	return bt.Root.Search(key)
}

func (bt *BTree) Insert(key string, value string) error {
	itemToInsert := blockUtil.NewItemWithKV(key, value)
	// if item already exists then just update key
	// todo: can keep a cache instead
	// may have issues when reconstructing from file
	if _, exists := bt.Root.Search(key); exists {
		return bt.Root.Update(*itemToInsert)
	}
	// if size of root is one away from max then we have to split it and create a new root
	if bt.Root.GetItemsSize() == consts.DefaultMinimumItems*2-1 {
		// generate new block ID
		generatedID, err := bt.Root.Bm.GenerateBlockID()
		if err != nil {
			return err
		}
		// Create a new block to replace root with id 0
		NewRoot := blockUtil.NewPnode(bt.Root.Bm).WithID(0)

		// set current root to new ID
		bt.Root.BlockID = generatedID

		//todo : can we save to disk at the end?
		if err = NewRoot.Commit(); err != nil {
			return err
		}
		if err = bt.Root.Commit(); err != nil {
			return err
		}

		// insert current root into new root's childs
		NewRoot.InsertChildAt(0, bt.Root)

		// split the old root as it is overflowing
		err = NewRoot.SplitChildAt(0)
		if err != nil {
			fmt.Printf("Unable to split the child: %s\n", err.Error())
			return err
		}

		// set the root of the Btree to be the new root
		bt.Root = NewRoot

		// find the child insertion index for item
		index, err := bt.Root.FindChildIndexForItem(key)
		if err != nil {
			fmt.Printf("Unable to get item insertion index: %s\n", err.Error())
			return err
		}

		// get the child node the item needs to be inserted into
		childNode, err := bt.Root.GetChildAt(index)
		if err != nil {
			fmt.Printf("Unable to get child node: %s\n", err.Error())
			return err
		}
		// insert child
		return childNode.InsertNonFull(itemToInsert)
	}
	return bt.Root.InsertNonFull(itemToInsert)
}
