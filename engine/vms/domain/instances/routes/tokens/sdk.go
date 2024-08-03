package tokens

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
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

// Adapter represents the tokens adapter
type Adapter interface {
	InstancesToBytes(ins Tokens) ([]byte, error)
	BytesToInstances(bytes []byte) (Tokens, error)
	InstanceToBytes(ins Token) ([]byte, error)
	BytesToInstance(bytes []byte) (Token, error)
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
