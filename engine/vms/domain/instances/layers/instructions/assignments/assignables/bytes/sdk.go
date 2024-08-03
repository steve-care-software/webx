package bytes

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

// NewBuilder creates a new bytes builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the bytes adapter
type Adapter interface {
	ToBytes(ins Bytes) ([]byte, error)
	ToInstance(bytes []byte) (Bytes, error)
}

// Builder represents a bytes builder
type Builder interface {
	Create() Builder
	WithJoin(join []string) Builder
	WithCompare(compare []string) Builder
	WithHashBytes(hashBytes string) Builder
	Now() (Bytes, error)
}

// Bytes represents the bytes assignable
type Bytes interface {
	Hash() hash.Hash
	IsJoin() bool
	Join() []string
	IsCompare() bool
	Compare() []string
	IsHashBytes() bool
	HashBytes() string
}
