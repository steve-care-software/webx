package sessions

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

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
}
