package block

import (
	"reflect"
	"testing"
)

func Test_block_creation(t *testing.T) {
	tests := []struct {
		name     string
		itemList []*Item
		children []uint64
	}{
		{
			name:     "empty-block",
			itemList: make([]*Item, 0),
			children: make([]uint64, 0),
		},
		{
			name: "filled",
			itemList: []*Item{
				NewItemWithKV("test-key", "test-value"),
			},
			children: []uint64{
				1,
			},
		},
	}
	for _, tt := range tests {
		block := NewBlock()
		block.setChildren(tt.children)
		block.setItemList(tt.itemList)

		if !reflect.DeepEqual(block.ChildrenIDs, tt.children) {
			t.Errorf("Children arrays are not equal")
		}
		if int(block.ItemListSize) != len(tt.itemList) {
			t.Errorf("Item list size is wrong")
		}

		if !reflect.DeepEqual(block.ItemList, tt.itemList) {
			t.Errorf("Children arrays are not equal")
		}
		if int(block.ChildrenSize) != len(tt.children) {
			t.Errorf("children size is wrong")
		}
	}
}

func Test_block_bytes_conversion(t *testing.T) {
	tests := []struct {
		name     string
		itemList []*Item
		children []uint64
	}{
		{
			name:     "empty-block",
			itemList: make([]*Item, 0),
			children: make([]uint64, 0),
		},
		{
			name: "filled",
			itemList: []*Item{
				NewItemWithKV("test-key", "test-value"),
			},
			children: []uint64{
				1,
			},
		},
	}
	for _, tt := range tests {
		block := NewBlock()
		block.setChildren(tt.children)
		block.setItemList(tt.itemList)

		buffer := block.IntoBytes()

		converted_block := NewBlock().FromBytes(buffer)
		if !reflect.DeepEqual(converted_block.ChildrenIDs, tt.children) {
			t.Errorf("Children arrays are not equal")
		}
		if int(converted_block.ItemListSize) != len(tt.itemList) {
			t.Errorf("Item list size is wrong")
		}

		if !reflect.DeepEqual(converted_block.ItemList, tt.itemList) {
			t.Errorf("Children arrays are not equal")
		}
		if int(converted_block.ChildrenSize) != len(tt.children) {
			t.Errorf("children size is wrong")
		}
	}
}
