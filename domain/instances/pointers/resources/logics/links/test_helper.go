package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
)

// NewLinksForTests creates a new links for tests
func NewLinksForTests(list []Link) Links {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkForTests creates new link for tests
func NewLinkForTests(elements elements.Elements) Link {
	ins, err := NewLinkBuilder().Create().
		WithElements(elements).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
