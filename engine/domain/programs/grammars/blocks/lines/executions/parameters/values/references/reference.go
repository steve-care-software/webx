package references

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"

type reference struct {
	element elements.Element
	index   uint
}

func createReference(
	element elements.Element,
	index uint,
) Reference {
	out := reference{
		element: element,
		index:   index,
	}

	return &out
}

// Element returns the element
func (obj *reference) Element() elements.Element {
	return obj.element
}

// Index returns the index
func (obj *reference) Index() uint {
	return obj.index
}
