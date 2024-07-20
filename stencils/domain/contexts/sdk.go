package contexts

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
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
	Head() hash.Hash
	Executions() []hash.Hash
}

// Repository represents a context repository
type Repository interface {
	Retrieve(dbPath []string) (Context, error)
}

// Service represents a service
type Service interface {
	Save(context Context) error
	Delete(hash hash.Hash) error
}
