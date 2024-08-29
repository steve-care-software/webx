package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type token struct {
	element     elements.Element
	cardinality cardinalities.Cardinality
	isReverse   bool
}

func createToken(
	element elements.Element,
	cardinality cardinalities.Cardinality,
	isReverse bool,
) Token {
	out := token{
		element:     element,
		cardinality: cardinality,
		isReverse:   isReverse,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.element.Name()
}

// Element returns the element
func (obj *token) Element() elements.Element {
	return obj.element
}

// Cardinality returns the cardinality
func (obj *token) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}

// IsReverse returns true if isReverse, false otherwise
func (obj *token) IsReverse() bool {
	return obj.isReverse
}
