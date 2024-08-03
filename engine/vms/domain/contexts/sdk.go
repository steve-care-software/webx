package contexts

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a context adapter
type Adapter interface {
	ToBytes(ins Context) ([]byte, error)
	ToInstance(bytes []byte) (Context, error)
}

// Builder represents a context builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier uint) Builder
	WithHead(head hash.Hash) Builder
	WithExecutions(executions []hash.Hash) Builder
	Now() (Context, error)
}

// Context represents a context
type Context interface {
	Hash() hash.Hash
	Identifier() uint
	Executions() []hash.Hash
	HasHead() bool
	Head() hash.Hash
}

// Repository represents a context repository
type Repository interface {
	Retrieve(dbPath []string) (Context, error)
}

// Service represents a service
type Service interface {
	Save(dbPath []string, context Context) error
	Delete(dbPath []string) error
}
