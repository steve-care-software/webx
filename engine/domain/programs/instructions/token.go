package instructions

type token struct {
	name     string
	elements Elements
	amount   uint
}

func createToken(
	name string,
	elements Elements,
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
func (obj *token) Elements() Elements {
	return obj.elements
}
