package compilers

import "github.com/steve-care-software/webx/selectors/domain/selectors"

type parameter struct {
	index    uint
	selector selectors.Selector
}

func createPrameter(
	index uint,
	selector selectors.Selector,
) Parameter {
	out := parameter{
		index:    index,
		selector: selector,
	}

	return &out
}

// Index returns the index
func (obj *parameter) Index() uint {
	return obj.index
}

// Selector returns the selector
func (obj *parameter) Selector() selectors.Selector {
	return obj.selector
}
