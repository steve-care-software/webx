package bytes

import (
	"encoding/binary"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
)

// AmountOfBytesIntUint64 represents the amount of bytes a uint64 contains
const AmountOfBytesIntUint64 = 8

// NewStateAdapter creates a new state adapter
func NewStateAdapter() states.Adapter {
	containerAdapter := NewContainerAdapter()
	builder := states.NewBuilder()
	stateBuilder := states.NewStateBuilder()
	return createStateAdapter(
		containerAdapter,
		builder,
		stateBuilder,
	)
}

// NewContainerAdapter creates a new container adapter
func NewContainerAdapter() containers.Adapter {
	pointerAdapter := NewPointerAdapter()
	builder := containers.NewBuilder()
	containerBuilder := containers.NewContainerBuilder()
	return createContainerAdapter(
		pointerAdapter,
		builder,
		containerBuilder,
	)
}

// NewPointerAdapter creates a new pointer adapter
func NewPointerAdapter() pointers.Adapter {
	retrievalAdapter := NewRetrievalAdapter()
	builder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	return createPointerAdapter(
		retrievalAdapter,
		builder,
		pointerBuilder,
	)
}

// NewRetrievalAdapter creates a new retrieval adapter
func NewRetrievalAdapter() retrievals.Adapter {
	builder := retrievals.NewBuilder()
	retrievalBuilder := retrievals.NewRetrievalBuilder()
	return createRetrievalAdapter(
		builder,
		retrievalBuilder,
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
