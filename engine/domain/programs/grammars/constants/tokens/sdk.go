package tokens

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/constants/tokens/elements"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// Builder represents a tokens builder
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents a list of tokens
type Tokens interface {
	List() []Token
	Fetch(name string) (Token, error)
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithElement(element elements.Element) TokenBuilder
	WithAmount(amount uint) TokenBuilder
	Now() (Token, error)
}

// Token represents a constant token
type Token interface {
	Name() string
	Element() elements.Element
	Amount() uint
}
