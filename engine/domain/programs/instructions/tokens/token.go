package tokens

import "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"

type token struct {
	name     string
	elements elements.Elements
	amount   uint
}

func createToken(
	name string,
	elements elements.Elements,
) Token {
	out := token{
		name:     name,
		elements: elements,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Elements returns the elements
func (obj *token) Elements() elements.Elements {
	return obj.elements
}
