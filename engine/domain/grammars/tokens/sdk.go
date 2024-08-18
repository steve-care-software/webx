package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// Builder represents a tokens list
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	List() []Token
	Fetch(name string) (Token, error)
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithElement(element elements.Element) TokenBuilder
	WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Element() elements.Element
	Cardinality() cardinalities.Cardinality
}
