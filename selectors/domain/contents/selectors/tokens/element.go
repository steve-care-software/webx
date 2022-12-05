package tokens

import "github.com/steve-care-software/webx/domain/databases/entities"

type element struct {
	el    entities.Identifier
	index uint
}

func createElement(
	el entities.Identifier,
	index uint,
) Element {
	out := element{
		el:    el,
		index: index,
	}

	return &out
}

// Element returns the element
func (obj *element) Element() entities.Identifier {
	return obj.el
}

// Index returns the index
func (obj *element) Index() uint {
	return obj.index
}
