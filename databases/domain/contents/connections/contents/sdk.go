package contents

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

const minContentSize = hash.Size + 8 + 1

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a connection adapter
type Adapter interface {
	ToContent(ins Content) ([]byte, error)
	ToInstance(content []byte) (Content, error)
}

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithData(data []byte) Builder
	WithKind(kind uint) Builder
	Now() (Content, error)
}

// Content represents a connection's content
type Content interface {
	Hash() hash.Hash
	Data() []byte
	Kind() uint
}
