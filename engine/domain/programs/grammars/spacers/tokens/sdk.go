package tokens

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/spacers/tokens/elements"

// Builder represents a tokens builder
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents a list of tokens
type Tokens interface {
	List() []Token
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithElement(element elements.Element) TokenBuilder
	WithAmount(amount uint) TokenBuilder
	Now() (Token, error)
}

// Token represents a spacer token
type Token interface {
	Name() string
	Element() elements.Element
	Amount() uint
}
