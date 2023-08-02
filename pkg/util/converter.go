package util

import "encoding/binary"

func Uint64ToBytes(index uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(index))
	return b
}

func Uint64FromBytes(b []byte) uint64 {
	return uint64(binary.LittleEndian.Uint64(b))
}

func Uint32ToBytes(val uint32) []byte {
	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, val)
	return buffer
}

func Uint32FromBytes(val []byte) uint32 {
	return uint32(binary.LittleEndian.Uint32(val))
}

func Uint16ToBytes(val uint16) []byte {
	buffer := make([]byte, 2)
	binary.LittleEndian.PutUint16(buffer, val)
	return buffer
}

func Uint16FromBytes(val []byte) uint16 {
	return uint16(binary.LittleEndian.Uint16(val))
}

func StrToBytes(val string) []byte {
	return []byte(val)
}

func StrFromBytes(val []byte) string {
	return string(val)
}
