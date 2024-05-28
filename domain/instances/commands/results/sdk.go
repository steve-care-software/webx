package results

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
)

const (
	// InputNotFoundError represents an input not found error
	InputNotFoundError (uint) = 1

	// InputNotBytesError represents an input not bytes error
	InputNotBytesError

	// OutputNotFoundError represents an output not found error
	OutputNotFoundError

	// OutputNotBytesError represents an output not bytes error
	OutputNotBytesError
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewSuccessBuilder creates a new success builder
func NewSuccessBuilder() SuccessBuilder {
	hashAdapter := hash.NewAdapter()
	return createSuccessBuilder(
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
	WithSuccess(success Success) Builder
	WithInterruption(interruption interruptions.Interruption) Builder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() Success
	IsInterruption() bool
	Interruption() interruptions.Interruption
}

// SuccessBuilder represents the success builder
type SuccessBuilder interface {
	Create() SuccessBuilder
	WithOutput(output outputs.Output) SuccessBuilder
	WithKind(kind kinds.Kind) SuccessBuilder
	Now() (Success, error)
}

// Success represents success result
type Success interface {
	Hash() hash.Hash
	Output() outputs.Output
	Kind() kinds.Kind
}
