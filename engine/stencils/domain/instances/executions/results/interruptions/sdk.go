package interruptions

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/interruptions/failures"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the interruption adapter
type Adapter interface {
	ToBytes(ins Interruption) ([]byte, error)
	ToInstance(bytes []byte) (Interruption, error)
}

// Builder represents an interruption builder
type Builder interface {
	Create() Builder
	WithStop(stopLine uint) Builder
	WithFailure(failure failures.Failure) Builder
	Now() (Interruption, error)
}

// Interruption represents an interruption
type Interruption interface {
	Hash() hash.Hash
	IsStop() bool
	Stop() *uint
	IsFailure() bool
	Failure() failures.Failure
}
