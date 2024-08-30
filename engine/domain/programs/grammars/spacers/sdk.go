package spacers

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/spacers/tokens"

// Builder represents the spacers builder
type Builder interface {
	Create() Builder
	WithList(list []Spacer) Builder
	Now() (Spacers, error)
}

// Spacers represents spacers
type Spacers interface {
	List() []Spacer
}

// SpacerBuilder represents the spacer builder
type SpacerBuilder interface {
	Create() SpacerBuilder
	WithName(name string) SpacerBuilder
	WithTokens(tokens tokens.Tokens) SpacerBuilder
	Now() (Spacer, error)
}

// Spacer represents a spacer
type Spacer interface {
	Name() string
	Tokens() tokens.Tokens
}
