package tokens

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

// NewBuilder creates a new token builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	hashAdapter := hash.NewAdapter()
	return createTokenBuilder(
		hashAdapter,
	)
}

// TokensAdapter represents the tokens adapter
type TokensAdapter interface {
	ToBytes(ins Tokens) ([]byte, error)
	ToInstance(bytes []byte) (Tokens, error)
}

// Builder represents tokens builder
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	Hash() hash.Hash
	List() []Token
}

// TokenAdapter represents the token adapter
type TokenAdapter interface {
	ToBytes(ins Token) ([]byte, error)
	ToInstance(bytes []byte) (Token, error)
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithElements(elements elements.Elements) TokenBuilder
	WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder
	WithOmission(omission omissions.Omission) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Elements() elements.Elements
	Cardinality() cardinalities.Cardinality
	HasOmission() bool
	Omission() omissions.Omission
}
