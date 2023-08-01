package btree

// consts "sixletters/kv-store/pkg/consts"

type Node struct {
	ItemList []*Item
	children []*Node
}

type BTree struct {
	Root     *Node
	minItems int
	maxItems int
}
