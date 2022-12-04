package channels

import "github.com/steve-care-software/webx/domain/cryptography/hash"

const minChannelSize = (hash.Size * 2) + 1

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

// Adapter represents a channel adapter
type Adapter interface {
	ToContent(ins Channel) ([]byte, error)
	ToChannel(content []byte) (Channel, error)
}

// Builder represents a channel builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithToken(token hash.Hash) Builder
	WithPrevious(previous hash.Hash) Builder
	WithNext(next hash.Hash) Builder
	Now() (Channel, error)
}

// Channel represents a chanel
type Channel interface {
	Hash() hash.Hash
	Token() hash.Hash
	HasPrevious() bool
	Previous() *hash.Hash
	HasNext() bool
	Next() *hash.Hash
}
