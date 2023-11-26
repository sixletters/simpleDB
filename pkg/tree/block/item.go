package block

import (
	"fmt"
	consts "sixletters/simple-db/pkg/consts"
	util "sixletters/simple-db/pkg/util"
)

type Item struct {
	keylen   uint16
	valueLen uint16
	key      string
	value    string
}

func (i *Item) setKey(key string) {
	i.key = key
	i.keylen = uint16(len(key))
}

func (i *Item) setValue(val string) {
	i.value = val
	i.valueLen = uint16(len(val))
}

func NewItemWithKV(key string, value string) *Item {
	item := new(Item)
	item.setKey(key)
	item.setValue(value)
	return item
}

func newItem() *Item {
	item := new(Item)
	item.setKey("")
	item.setValue("")
	return item
}

func (i *Item) validate() error {
	if i.keylen > consts.DefaultItemMaxKeyLen {
		return fmt.Errorf("key length exceeds max keylength of %d", i.keylen)
	}
	if i.valueLen > consts.DefaultItemMaxValueLen {
		return fmt.Errorf("key length exceeds max keylength of %d", i.valueLen)
	}
	return nil
}

func (i *Item) IntoBytes() []byte {
	byteBuffer := make([]byte, consts.ItemSize)
	offset := uint16(0)
	copy(byteBuffer[offset:], util.Uint16ToBytes(i.keylen))
	offset += 2

	copy(byteBuffer[offset:], util.Uint16ToBytes(i.valueLen))
	offset += 2

	copy(byteBuffer[offset:], util.StrToBytes(i.key))
	offset += i.keylen

	copy(byteBuffer[offset:], util.StrToBytes(i.value))
	return byteBuffer
}

func (i *Item) FromBytes(bytebuffer []byte) *Item {
	offset := uint16(0)
	i.keylen = util.Uint16FromBytes(bytebuffer[offset : offset+2])
	offset += 2

	i.valueLen = util.Uint16FromBytes(bytebuffer[offset : offset+2])
	offset += 2

	i.key = util.StrFromBytes(bytebuffer[offset : offset+i.keylen])
	offset += i.keylen

	i.value = util.StrFromBytes(bytebuffer[offset : offset+i.valueLen])
	return i
}
