package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/references"
)

// NewLinkForTests creates new link for tests
func NewLinkForTests(elements elements.Elements) Link {
	ins, err := NewBuilder().Create().
		WithElements(elements).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithReferencesForTests creates a new link with references for tests
func NewLinkWithReferencesForTests(elements elements.Elements, references references.References) Link {
	ins, err := NewBuilder().Create().
		WithElements(elements).
		WithReferences(references).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
