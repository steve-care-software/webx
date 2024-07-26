package files

import "encoding/binary"

func bytesToUint64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}

func uint64ToBytes(value uint64) []byte {
	bytes := make([]byte, amountOfBytesIntUint64)
	binary.BigEndian.PutUint64(bytes, value)
	return bytes
}
