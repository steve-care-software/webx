package failures

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new failure builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the failure builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithCode(code uint) Builder
	IsRaisedInLayer() Builder
	WithMessage(message string) Builder
	Now() (Failure, error)
}

// Failure represents failure result
type Failure interface {
	Hash() hash.Hash
	Index() uint
	Code() uint
	IsRaisedInLayer() bool
	HasMessage() bool
	Message() string
}
