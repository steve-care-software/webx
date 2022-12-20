package histories

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

const historySize = hash.Size + 8

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an histories adapter
type Adapter interface {
	ToContent(ins History) ([]byte, error)
	ToHistory(content []byte) (History, error)
}

// Builder represents an history builder
type Builder interface {
	Create() Builder
	WithCommit(commit hash.Hash) Builder
	WithScore(score uint) Builder
	Now() (History, error)
}

// History represents a commit history
type History interface {
	Commit() hash.Hash
	Score() uint
}
