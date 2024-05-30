package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
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
