package btree

import (
	"fmt"
	"sixletters/simple-db/pkg/consts"
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
	BlockID  uint64
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
		fmt.Printf("Unable to get children of current node with id: %d", pn.BlockID)
	}
	for index, child := range childrenList {
		fmt.Println("Printing ", index+1, " th child")
		child.PrintTree()
	}
}

// Function converts a Pnode from a serialized block to a pnode
func (pn *Pnode) FromBlock(block *Block) *Pnode {
	pn.BlockID = block.Id
	// TODO: May need to deep copy in the future.
	pn.children = block.ChildrenIDs
	pn.Items = block.ItemList
	return pn
}

// Function converts a pnode into a serialized block
func (pn *Pnode) IntoBlock() *Block {
	block := NewBlock()
	block.Id = pn.BlockID
	// TODO: May need to deep copy in the future.
	block.setChildren(pn.children)
	block.setItemList(pn.Items)
	return block
}

func (pn *Pnode) IsLeaf() bool {
	return len(pn.children) == 0
}

// Returns the items in a Pnode
func (pn *Pnode) GetItems() []*Item {
	return pn.Items
}

func (pn *Pnode) WithItems(itemList []*Item) *Pnode {
	pn.Items = itemList
	return pn
}

func (pn *Pnode) WithID(ID uint64) *Pnode {
	pn.BlockID = ID
	return pn
}

func (pn *Pnode) WithChildren(children []uint64) *Pnode {
	pn.children = children
	return pn
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

func (pn *Pnode) GetItemsSize() int {
	return len(pn.Items)
}

func (pn *Pnode) GetChildrenSize() int {
	return len(pn.children)
}

// Function returns the child Pnode at a particular index, throws an error if index is invalid
func (pn *Pnode) GetChildAt(index int) (*Pnode, error) {
	if index >= len(pn.children) {
		return nil, fmt.Errorf("index is bigger than the number of children in this node")
	}
	// retrieve corresponding block by blockID using block manager
	block, err := pn.Bm.GetBlockByID(int64(pn.children[index]))
	if err != nil {
		return nil, err
	}
	//Deserialize to pnode
	return NewPnode(pn.Bm).FromBlock(block), nil
}

// Functions returns all child Pnodes of a Pnode
func (pn *Pnode) GetChildPnodes() ([]*Pnode, error) {
	childPnodes := make([]*Pnode, len(pn.children))
	for i, childBlockID := range pn.children {
		// retrieve corresponding block by blockID using block manager
		block, err := pn.Bm.GetBlockByID(int64(childBlockID))
		if err != nil {
			return nil, err
		}
		// deserialize to Pnode
		childPnodes[i] = NewPnode(pn.Bm).FromBlock(block)
	}
	return childPnodes, nil
}

// Returns an item at a particular index of a Pnode, returns an error if index is invalid/
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
// Todo: Could be optimized to use binary search next time
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

func (pn *Pnode) FindChildIndexForItem(key string) (int, error) {
	for index, item := range pn.GetItems() {
		if key < item.key {
			return index, nil
		}
	}
	return len(pn.children) - 1, nil
}

func (pn *Pnode) GetLastChild() (*Pnode, error) {
	return pn.GetChildAt(len(pn.children) - 1)
}
func (pn *Pnode) SearchCurrentNode(key string) (string, bool) {
	for _, item := range pn.Items {
		if item.key == key {
			return item.value, true
		}
	}
	return "", false
}

func (pn *Pnode) Search(key string) (string, bool) {
	value, found := pn.SearchCurrentNode(key)
	if found {
		return value, true
	}
	if pn.IsLeaf() {
		return "", false
	}

	childIndex, err := pn.FindChildIndexForItem(key)
	if err != nil {
		return "", false
	}

	child, err := pn.GetChildAt(childIndex)
	if err != nil {
		return "", false
	}
	return child.Search(key)
}

func (pn *Pnode) GetValue(key string) (string, bool) {
	return pn.Search(key)
}

func (pn *Pnode) SavetoDisk() error {
	block := pn.IntoBlock()
	return pn.Bm.WriteBlockToDisk(block)
}

func (pn *Pnode) SplitNode() (*Item, *Pnode, *Pnode, error) {
	if pn.IsLeaf() {
		return pn.splitLeafNode()
	}
	return pn.splitNodeWithChildren()
}

func (pn *Pnode) SplitChildAt(index int) error {
	if index >= len(pn.children) {
		return fmt.Errorf("there is no child at this index")
	}
	childBlockID := pn.children[index]
	childBlock, err := pn.Bm.GetBlockByID(int64(childBlockID))
	if err != nil {
		return fmt.Errorf("unable to get child Block")
	}

	childNode := NewPnode(pn.Bm).FromBlock(childBlock)
	item, left, right, err := childNode.SplitNode()
	if err != nil {
		return fmt.Errorf("unable to split child node")
	}
	pn.Items = util.InsertAt(pn.Items, item, index)
	pn.children[index] = left.BlockID
	pn.children = util.InsertAt(pn.children, right.BlockID, index+1)
	pn.SavetoDisk()
	return nil
}

func (pn *Pnode) splitNodeWithChildren() (*Item, *Pnode, *Pnode, error) {
	if pn.IsLeaf() {
		return nil, nil, nil, fmt.Errorf("unable to split a leaf node as if it had children")
	}

	items := pn.GetItems()
	midIndex := (len(items) / 2) + 1
	midItem := items[midIndex]
	pn.Items = items[0:midIndex]
	SplitNodeItems := items[midIndex+1:]

	children := pn.GetChildren()
	pn.children = children[0 : midIndex+1]
	SplitNodeChildren := children[midIndex+1:]

	generatedID, err := pn.Bm.GenerateBlockID()
	if err != nil {
		return nil, nil, nil, err
	}

	err = pn.SavetoDisk()
	if err != nil {
		return nil, nil, nil, err
	}

	SplitNode := NewPnode(pn.Bm).WithItems(SplitNodeItems).WithChildren(SplitNodeChildren).WithID(generatedID)
	err = SplitNode.SavetoDisk()
	if err != nil {
		return nil, nil, nil, err
	}

	return midItem, pn, SplitNode, nil
}

func (pn *Pnode) splitLeafNode() (*Item, *Pnode, *Pnode, error) {
	if !pn.IsLeaf() {
		return nil, nil, nil, fmt.Errorf("unable to split a leaf node that is not a leaf node")
	}

	items := pn.GetItems()
	midIndex := (len(items) / 2) + 1
	midItem := items[midIndex]
	pn.Items = items[0:midIndex]
	SplitNodeItems := items[midIndex+1:]

	generatedID, err := pn.Bm.GenerateBlockID()
	if err != nil {
		return nil, nil, nil, err
	}

	err = pn.SavetoDisk()
	if err != nil {
		return nil, nil, nil, err
	}

	SplitNode := NewPnode(pn.Bm).WithItems(SplitNodeItems).WithID(generatedID)
	err = SplitNode.SavetoDisk()
	if err != nil {
		return nil, nil, nil, err
	}

	return midItem, pn, SplitNode, nil
}

func (pn *Pnode) SetChildAt(index int, child *Pnode) error {
	if index > len(pn.children) {
		return fmt.Errorf("out of range error for child index setting")
	}
	pn.children[index] = child.BlockID
	return nil
}

func (pn *Pnode) InsertChildAt(index int, child *Pnode) error {
	if index > len(pn.children) {
		return fmt.Errorf("out of range error for child index setting")
	}
	pn.children = util.InsertAt(pn.children, child.BlockID, index)
	return nil
}

func (pn *Pnode) AddItemAndChildren(item *Item, leftNode *Pnode, rightnode *Pnode) error {
	insertionIndex := pn.AddItem(item)
	err := pn.SetChildAt(insertionIndex, leftNode)
	if err != nil {
		return fmt.Errorf("unable to insert child at %d", insertionIndex)
	}

	err = pn.InsertChildAt(insertionIndex+1, rightnode)
	if err != nil {
		return fmt.Errorf("unable to insert child at %d", insertionIndex+1)
	}
	return nil
}

func (pn *Pnode) InsertNonFull(item *Item) error {
	// If leaf add item, save to disk and done.
	if pn.IsLeaf() {
		pn.AddItem(item)
		return pn.SavetoDisk()
	}
	// Not leaf node

	// Find corresponding child index to be inserted to
	childIndex, err := pn.FindChildIndexForItem(item.key)
	if err != nil {
		return fmt.Errorf("unable to find child for this key")
	}
	// Retrieve the child node
	ChildNode, err := pn.GetChildAt(childIndex)
	if err != nil {
		return fmt.Errorf("unable to find child node")
	}

	// Check if child node is 1 below the limit, if it is then split the node
	if ChildNode.GetItemsSize() == consts.DefaultMinimumItems*2-1 {
		pn.SplitChildAt(childIndex)
		// If the newly inserted item from the split is smaller than item, then increase insertion index as
		// it should be inserted to tree on the right of newly inserted item from the split.
		if item.key > pn.GetItems()[childIndex].key {
			childIndex += 1
		}
	}

	// retrieve child
	ChildNode, err = pn.GetChildAt(childIndex)
	if err != nil {
		return fmt.Errorf("unable to find child node")
	}
	return ChildNode.InsertNonFull(item)
}
