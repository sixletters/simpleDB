package btree

import (
	"encoding/binary"
	"fmt"
	consts "sixletters/kv-store/pkg/consts"
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
	i.key = val
	i.keylen = uint16(len(val))
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
		return fmt.Errorf("Key length exceeds max keylength of %d", i.keylen)
	}
	if i.valueLen > consts.DefaultItemMaxValueLen {
		return fmt.Errorf("Key length exceeds max keylength of %d", i.valueLen)
	}
	return nil
}

func uint16ToBytes(val uint16) []byte {
	buffer := make([]byte, 2)
	binary.LittleEndian.PutUint16(buffer, val)
	return buffer
}

func uint16FromBytes(val []byte) uint16 {
	return uint16(binary.LittleEndian.Uint16(val))
}

func strToBytes(val string) []byte {
	return []byte(val)
}

func strFromBytes(val []byte) string {
	return string(val)
}

func (i *Item) intoBytes() []byte {
	byteBuffer := make([]byte, consts.DefaultItemMaxSize)
	offset := uint16(0)
	copy(byteBuffer[offset:], uint16ToBytes(i.keylen))
	offset += 2

	copy(byteBuffer[offset:], uint16ToBytes(i.valueLen))
	offset += 2

	copy(byteBuffer[offset:], strToBytes(i.key))
	offset += i.keylen

	copy(byteBuffer[offset:], strToBytes(i.value))
	return byteBuffer
}

func (i *Item) validateByteBuffer() {

}

func (i *Item) fromBytes(bytebuffer []byte) *Item {
	offset := uint16(0)
	i.keylen = uint16FromBytes(bytebuffer[offset:2])
	offset += 2

	i.valueLen = uint16FromBytes(bytebuffer[offset:2])
	offset += 2

	i.key = strFromBytes(bytebuffer[offset : offset+i.keylen])
	offset += i.keylen

	i.value = strFromBytes(bytebuffer[offset : offset+i.valueLen])
	return i
}
