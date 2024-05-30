package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/references"
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

// NewLinkWithReferecesForTests creates new link with references for tests
func NewLinkWithReferecesForTests(elements elements.Elements, references references.References) Link {
	ins, err := NewLinkBuilder().Create().
		WithElements(elements).
		WithReferences(references).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
