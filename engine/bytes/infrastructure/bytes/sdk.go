package bytes

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
	"github.com/steve-care-software/webx/engine/bytes/domain/pointers"
)

const (
	// DelimiterFlag represents the delimiter flag
	DelimiterFlag uint8 = iota

	// represents the pointer flag
	PointerFlag

	// NamespaceFlag represents the namespace flag
	NamespaceFlag
)

// AmountOfBytesIntUint64 represents the amount of bytes a uint64 contains
const AmountOfBytesIntUint64 = 8

const invalidDataPatternErr = "the data was expected to contain at least %d bytes, %d provided"

// NewNamespaceAdapter creates a new namespace adapter
func NewNamespaceAdapter() namespaces.Adapter {
	delimiterAdapter := NewDelimiterAdapter()
	builder := namespaces.NewBuilder()
	namespaceBuilder := namespaces.NewNamespaceBuilder()
	return createNamespaceAdapter(
		delimiterAdapter,
		builder,
		namespaceBuilder,
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

// AmountReturnRemaining returns the amount of return elements remaining
func AmountReturnRemaining(data []byte) (int, []byte, error) {
	expectation := AmountOfBytesIntUint64
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return 0, nil, errors.New(str)
	}

	return int(BytesToUint64(data[:expectation])), data[expectation:], nil
}
