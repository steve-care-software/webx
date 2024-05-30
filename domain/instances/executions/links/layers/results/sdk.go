package results

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/success"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the result adapter
type Adapter interface {
	ToBytes(ins Result) ([]byte, error)
	ToInstance(bytes []byte) (Result, error)
}

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithSuccess(success success.Success) Builder
	WithInterruption(interruption interruptions.Interruption) Builder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() success.Success
	IsInterruption() bool
	Interruption() interruptions.Interruption
}
