package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/references"
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
	ins, err := NewLinkBuilder().Create().
		WithOrigin(origin).
		WithElements(elements).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithReferecesForTests creates new link with references for tests
func NewLinkWithReferecesForTests(origin origins.Origin, elements elements.Elements, references references.References) Link {
	ins, err := NewLinkBuilder().Create().
		WithOrigin(origin).
		WithElements(elements).
		WithReferences(references).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
