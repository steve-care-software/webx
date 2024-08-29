package tokens

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/reverses"
)

type token struct {
	element     elements.Element
	cardinality cardinalities.Cardinality
	reverse     reverses.Reverse
}

func createToken(
	element elements.Element,
	cardinality cardinalities.Cardinality,
) Token {
	return createTokenInternally(element, cardinality, nil)
}

func createTokenWithReverse(
	element elements.Element,
	cardinality cardinalities.Cardinality,
	reverse reverses.Reverse,
) Token {
	return createTokenInternally(element, cardinality, reverse)
}

func createTokenInternally(
	element elements.Element,
	cardinality cardinalities.Cardinality,
	reverse reverses.Reverse,
) Token {
	out := token{
		element:     element,
		cardinality: cardinality,
		reverse:     reverse,
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

// HasReverse returns true if there is a reverse, false otherwise
func (obj *token) HasReverse() bool {
	return obj.reverse != nil
}

// Reverse returns the reverse, if any
func (obj *token) Reverse() reverses.Reverse {
	return obj.reverse
}
