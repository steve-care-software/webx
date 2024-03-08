package kinds

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewKindBuilder creates a new kind builder
func NewKindBuilder() KindBuilder {
	hashAdapter := hash.NewAdapter()
	return createKindBuilder(
		hashAdapter,
	)
}

// KindBuilder represents a kind builder
type KindBuilder interface {
	Create() KindBuilder
	IsPrompt() KindBuilder
	IsContinue() KindBuilder
	Now() (Kind, error)
}

// Kind represents the output kind
type Kind interface {
	Hash() hash.Hash
	IsPrompt() bool
	IsContinue() bool
}
