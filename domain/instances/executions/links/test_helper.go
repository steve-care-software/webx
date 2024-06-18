package links

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links"
)

// NewLinkForTests creates a new link for tests
func NewLinkForTests(input []byte, source links.Link) Link {
	ins, err := NewBuilder().Create().
		WithInput(input).
		WithSource(source).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkWithLayersForTests creates a new link with layers for tests
func NewLinkWithLayersForTests(input []byte, source links.Link, layers layers.Layers) Link {
	ins, err := NewBuilder().Create().
		WithInput(input).
		WithSource(source).
		WithLayers(layers).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
