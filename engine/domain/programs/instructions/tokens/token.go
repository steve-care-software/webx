package tokens

import "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"

type token struct {
	name    string
	element elements.Element
	amount  uint
}

func createToken(
	name string,
	element elements.Element,
	amount uint,
) Token {
	out := token{
		name:    name,
		element: element,
		amount:  amount,
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

// Amount returns the amount
func (obj *token) Amount() uint {
	return obj.amount
}
