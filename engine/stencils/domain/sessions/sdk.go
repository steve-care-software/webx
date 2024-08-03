package sessions

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a session adapter
type Adapter interface {
	ToBytes(ins Session) ([]byte, error)
	ToInstance(bytes []byte) (Session, error)
}

// Builder represents a session builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithExecutions(executions []executions.Executions) Builder
	Now() (Session, error)
}

// Session represents a session
type Session interface {
	Hash() hash.Hash
	Executions() []executions.Executions
	Token() hash.Hash
}
