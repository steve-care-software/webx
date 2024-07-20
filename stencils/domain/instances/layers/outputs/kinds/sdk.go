package kinds

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

// NewBuilder creates a new kind builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the kind adapter
type Adapter interface {
	ToBytes(ins Kind) ([]byte, error)
	ToInstance(bytes []byte) (Kind, error)
}

// Builder represents a kind builder
type Builder interface {
	Create() Builder
	IsPrompt() Builder
	IsContinue() Builder
	Now() (Kind, error)
}

// Kind represents the output kind
type Kind interface {
	Hash() hash.Hash
	IsPrompt() bool
	IsContinue() bool
}
