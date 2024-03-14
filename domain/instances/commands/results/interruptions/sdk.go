package interruptions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions/failures"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
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
