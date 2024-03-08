package links

import (
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
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
func NewLinkForTests(origin origins.Origin, elements elements.Elements) Link {
	ins, err := NewLinkBuilder().Create().WithOrigin(origin).WithElements(elements).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
