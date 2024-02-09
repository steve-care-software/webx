package results

import (
	"github.com/steve-care-software/identity/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
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

// NewFailureBuilder creates a new failure builder
func NewFailureBuilder() FailureBuilder {
	hashAdapter := hash.NewAdapter()
	return createFailureBuilder(
		hashAdapter,
	)
}

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithSuccess(success Success) Builder
	WithFailure(failure Failure) Builder
	Now() (Result, error)
}

// Result represents result
type Result interface {
	Hash() hash.Hash
	IsSuccess() bool
	Success() Success
	IsFailure() bool
	Failure() Failure
}

// SuccessBuilder represents the success builder
type SuccessBuilder interface {
	Create() SuccessBuilder
	WithBytes(bytes []byte) SuccessBuilder
	WithKind(kind layers.Kind) SuccessBuilder
	Now() (Success, error)
}

// Success represents success result
type Success interface {
	Hash() hash.Hash
	Bytes() []byte
	Kind() layers.Kind
}

// FailureBuilder represents the failure builder
type FailureBuilder interface {
	Create() FailureBuilder
	WithCode(code uint) FailureBuilder
	IsRaisedInLayer() FailureBuilder
	Now() (Failure, error)
}

// Failure represents failure result
type Failure interface {
	Hash() hash.Hash
	Code() uint
	IsRaisedInLayer() bool
}
