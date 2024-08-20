package uints

import "encoding/binary"

func uintToBytes(value uint64) []byte {
	bytes := make([]byte, AmountOfBytesIntUint64)
	binary.BigEndian.PutUint64(bytes, uint64(value))
	return bytes
}

func bytesToUInt(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}
