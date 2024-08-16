package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"
)

// Builder represents the token builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithElement(element elements.Element) Builder
	WithCardinality(cardinality cardinalities.Cardinality) Builder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Element() elements.Element
	Cardinality() cardinalities.Cardinality
}
