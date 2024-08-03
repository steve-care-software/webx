package failures

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

// NewBuilder creates a new failure builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the failure adapter
type Adapter interface {
	ToBytes(ins Failure) ([]byte, error)
	ToInstance(bytes []byte) (Failure, error)
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
