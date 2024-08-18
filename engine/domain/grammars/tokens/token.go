package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"
)

type token struct {
	element     elements.Element
	cardinality cardinalities.Cardinality
}

func createToken(
	element elements.Element,
	cardinality cardinalities.Cardinality,
) Token {
	out := token{
		element:     element,
		cardinality: cardinality,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	if obj.element.IsBlock() {
		return obj.element.Block()
	}

	return obj.element.Rule()
}

// Element returns the element
func (obj *token) Element() elements.Element {
	return obj.element
}

// Cardinality returns the cardinality
func (obj *token) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}
