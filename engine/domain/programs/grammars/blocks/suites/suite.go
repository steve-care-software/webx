package suites

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"

type suite struct {
	name    string
	element elements.Element
	isFail  bool
}

func createSuite(
	name string,
	element elements.Element,
	isFail bool,
) Suite {
	out := suite{
		name:    name,
		element: element,
		isFail:  isFail,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Element returns the element
func (obj *suite) Element() elements.Element {
	return obj.element
}

// IsFail returns true if expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}
