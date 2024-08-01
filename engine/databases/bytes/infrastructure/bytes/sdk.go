package bytes

import (
	"encoding/binary"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

// AmountOfBytesIntUint64 represents the amount of bytes a uint64 contains
const AmountOfBytesIntUint64 = 8

// NewStateAdapter creates a new state adapter
func NewStateAdapter() states.Adapter {
	pointerAdapter := NewPointerAdapter()
	builder := states.NewBuilder()
	stateBuilder := states.NewStateBuilder()
	return createStateAdapter(
		pointerAdapter,
		builder,
		stateBuilder,
	)
}

// NewPointerAdapter creates a new pointer adapter
func NewPointerAdapter() pointers.Adapter {
	delimiterAdapter := NewDelimiterAdapter()
	builder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	return createPointerAdapter(
		delimiterAdapter,
		builder,
		pointerBuilder,
	)
}

// NewDelimiterAdapter creates a new delimiter adapter
func NewDelimiterAdapter() delimiters.Adapter {
	builder := delimiters.NewBuilder()
	delimiterBuilder := delimiters.NewDelimiterBuilder()
	return createDelimiterAdapter(
		builder,
		delimiterBuilder,
	)
}

// BytesToUint64 converts bytes to uint64
func BytesToUint64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}

// Uint64ToBytes converts uint64 to bytes
func Uint64ToBytes(value uint64) []byte {
	bytes := make([]byte, AmountOfBytesIntUint64)
	binary.BigEndian.PutUint64(bytes, value)
	return bytes
}
