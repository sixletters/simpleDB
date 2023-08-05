package btree

import (
	"fmt"
	"sixletters/simple-db/pkg/util"
)

type Node interface {
	insert(key string, value string) error
	get(key string) (string, error)
	printTree(level int)
}

type Pnode struct {
	Items    []*Item
	children []uint64
	ID       uint64
	Bm       *BlockManager
}

func NewPnode(bm *BlockManager) *Pnode {
	pnode := &Pnode{}
	pnode.Items = make([]*Item, 0)
	pnode.children = make([]uint64, 0)
	pnode.Bm = bm
	return pnode
}

func (pn *Pnode) PrintPnode() {
	fmt.Println("Pnode")
	fmt.Println("----------")
	for _, item := range pn.Items {
		fmt.Println(*item)
	}
	fmt.Println("***********************")
}

func (pn *Pnode) PrintTree() {
	pn.PrintPnode()
	childrenList, err := pn.GetChildPnodes()
	if err != nil {
		fmt.Printf("Unable to get children of current node with id: %d", pn.ID)
	}
	for index, child := range childrenList {
		fmt.Println("Printing ", index+1, " th child")
		child.PrintTree()
	}
}

func (pn *Pnode) FromBlock(block *Block) *Pnode {
	pn.ID = block.Id
	// TODO: May need to deep copy in the future.
	pn.children = block.ChildrenIDs
	pn.Items = block.ItemList
	return pn
}

func (pn *Pnode) IntoBlock() *Block {
	block := NewBlock()
	block.Id = pn.ID
	// TODO: May need to deep copy in the future.
	block.setChildren(pn.children)
	block.setItemList(pn.Items)
	return block
}

func (pn *Pnode) IsLeaf() bool {
	return len(pn.children) == 0
}

func (pn *Pnode) GetItems() []*Item {
	return pn.Items
}

func (pn *Pnode) SetItems(itemList []*Item) {
	pn.Items = itemList
}

func (pn *Pnode) GetChildren() []uint64 {
	return pn.children
}

func (pn *Pnode) SetChildren(children []uint64) {
	pn.children = children
}

func (pn *Pnode) GetChildAt(index int) (*Pnode, error) {
	if index >= len(pn.children) {
		return nil, fmt.Errorf("index is bigger than the number of children in this node")
	}
	block, err := pn.Bm.GetBlockByID(int64(pn.children[index]))
	if err != nil {
		return nil, err
	}
	return NewPnode(pn.Bm).FromBlock(block), nil
}

func (pn *Pnode) GetChildPnodes() ([]*Pnode, error) {
	childPnodes := make([]*Pnode, len(pn.children))
	for i, childBlockID := range pn.children {
		block, err := pn.Bm.GetBlockByID(int64(childBlockID))
		if err != nil {
			return nil, err
		}
		childPnodes[i] = NewPnode(pn.Bm).FromBlock(block)
	}
	return childPnodes, nil
}

func (pn *Pnode) GetItemAt(index int) (*Item, error) {
	if index >= len(pn.Items) {
		return nil, fmt.Errorf("index is bigger than the number of items in node")
	}
	return pn.Items[index], nil
}

func (pn *Pnode) InsertItemAt(index int, item *Item) error {
	if index > len(pn.Items) {
		return fmt.Errorf("the index is bigger than the amount of elements in the item list")
	}
	pn.Items = util.InsertAt(pn.GetItems(), item, index)
	return nil
}

func (pn *Pnode) IsOverflown() bool {
	return len(pn.Items) > pn.Bm.GetMaxItemsSize()
}

// adds an element and returns the index -> adds via a linear search
// Could be optimized to use binary search next time
func (pn *Pnode) AddItem(itemToAdd *Item) int {
	for index, item := range pn.Items {
		if itemToAdd.key > item.key {
			continue
		}
		pn.Items = util.InsertAt(pn.Items, itemToAdd, index)
		return index
	}

	// will only reach here if the key is bigger than all elements in the item list
	LastIndex := len(pn.Items)
	pn.Items = util.InsertAt(pn.Items, itemToAdd, LastIndex)
	return LastIndex
}
