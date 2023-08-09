package btree

import (
	"fmt"
	"os"
	"sixletters/simple-db/pkg/consts"
)

// consts "sixletters/kv-store/pkg/consts"

type BTree struct {
	Root     *Pnode
	minItems int
}

func NewBTree(file *os.File) *BTree {
	BlockManager := NewBlockManager(file)
	tree := &BTree{}
	tree.Root = NewPnode(BlockManager)
	tree.minItems = consts.DefaultMinimumItems
	return tree
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
	itemToInsert := NewItemWithKV(key, value)
	if bt.Root.GetItemsSize() == consts.DefaultMinimumItems*2-1 {
		generatedID, err := bt.Root.Bm.GenerateBlockID()
		if err != nil {
			return err
		}

		NewRoot := NewPnode(bt.Root.Bm).WithID(0)
		bt.Root.BlockID = generatedID
		NewRoot.SavetoDisk()
		bt.Root.SavetoDisk()

		NewRoot.InsertChildAt(0, bt.Root)
		err = NewRoot.SplitChildAt(0)
		if err != nil {
			fmt.Printf("Unable to split the child")
		}

		bt.Root = NewRoot
		index, err := bt.Root.FindChildIndexForItem(key)
		if err != nil {
			fmt.Printf("Unable to get item insertion index")
		}

		childNode, err := bt.Root.GetChildAt(index)
		if err != nil {
			fmt.Printf("Unable to get child node")
		}
		return childNode.InsertNonFull(itemToInsert)
	}
	return bt.Root.InsertNonFull(itemToInsert)
}
