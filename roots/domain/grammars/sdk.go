package grammars

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

const minGrammarSize = hash.Size * 2

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

// Adapter represents a grammar adapter
type Adapter interface {
	ToContent(ins Grammar) ([]byte, error)
	ToGrammar(content []byte) (Grammar, error)
}

// Builder represents a grammar
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithRoot(root hash.Hash) Builder
	WithChannels(channels []hash.Hash) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Hash() hash.Hash
	Root() hash.Hash
	HasChannels() bool
	Channels() []hash.Hash
}
