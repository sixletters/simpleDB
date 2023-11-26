package block

import (
	"sixletters/simple-db/pkg/consts"
	"sixletters/simple-db/pkg/util"
)

// Total size is 4096
type Block struct {
	Id           uint64   // 8
	ChildrenSize uint32   // 4
	ItemListSize uint32   // 4
	ChildrenIDs  []uint64 // max is 8 * 30 = 240
	ItemList     []*Item  //  max is 128 * 30 = 3840
}

func (b *Block) setItemList(itemList []*Item) {
	b.ItemList = itemList
	b.ItemListSize = uint32(len(itemList))
}

func (b *Block) setChildren(childrenIDs []uint64) {
	b.ChildrenIDs = childrenIDs
	b.ChildrenSize = uint32(len(childrenIDs))
}

func NewBlock() *Block {
	block := &Block{}
	block.ChildrenIDs = make([]uint64, 0)
	block.ItemList = make([]*Item, 0)
	return block
}

func (b *Block) IntoBytes() []byte {
	buffer := make([]byte, consts.BlockSize)
	offset := 0

	copy(buffer[offset:offset+8], util.Uint64ToBytes(b.Id))
	offset += 8

	copy(buffer[offset:offset+4], util.Uint32ToBytes(b.ChildrenSize))
	offset += 4

	copy(buffer[offset:offset+4], util.Uint32ToBytes(b.ItemListSize))
	offset += 4

	for _, id := range b.ChildrenIDs {
		copy(buffer[offset:offset+8], util.Uint64ToBytes(id))
		offset += 8
	}

	for _, item := range b.ItemList {
		copy(buffer[offset:offset+consts.ItemSize], item.IntoBytes())
		offset += consts.ItemSize
	}

	return buffer
}

func (b *Block) FromBytes(buffer []byte) *Block {
	offset := 0
	b.Id = util.Uint64FromBytes(buffer[offset : offset+8])
	offset += 8

	b.ChildrenSize = util.Uint32FromBytes(buffer[offset : offset+4])
	offset += 4

	b.ItemListSize = util.Uint32FromBytes(buffer[offset : offset+4])
	offset += 4

	b.ChildrenIDs = make([]uint64, b.ChildrenSize)
	for i := 0; i < int(b.ChildrenSize); i++ {
		b.ChildrenIDs[i] = util.Uint64FromBytes(buffer[offset : offset+8])
		offset += 8
	}

	b.ItemList = make([]*Item, b.ItemListSize)
	for i := 0; i < int(b.ItemListSize); i++ {
		b.ItemList[i] = newItem().FromBytes(buffer[offset : offset+consts.ItemSize])
		offset += consts.ItemSize
	}

	return b
}
