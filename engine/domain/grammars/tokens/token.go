package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"
)

type token struct {
	name        string
	element     elements.Element
	cardinality cardinalities.Cardinality
}

func createToken(
	name string,
	element elements.Element,
	cardinality cardinalities.Cardinality,
) Token {
	out := token{
		name:        name,
		element:     element,
		cardinality: cardinality,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Element returns the element
func (obj *token) Element() elements.Element {
	return obj.element
}

// Cardinality returns the cardinality
func (obj *token) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}
