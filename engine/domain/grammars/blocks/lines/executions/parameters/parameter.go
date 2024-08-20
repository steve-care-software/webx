package parameters

import "github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"

type parameter struct {
	element elements.Element
	index   uint
	name    string
}

func createParameter(
	element elements.Element,
	index uint,
	name string,
) Parameter {
	out := parameter{
		element: element,
		index:   index,
		name:    name,
	}

	return &out
}

// Element returns the element
func (obj *parameter) Element() elements.Element {
	return obj.element
}

// Index returns the index
func (obj *parameter) Index() uint {
	return obj.index
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}
