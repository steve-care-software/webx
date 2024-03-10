package libraries

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links"
)

// NewLibraryForTests creates a new library for tests
func NewLibraryForTests(layers layers.Layers) Library {
	ins, err := NewBuilder().Create().WithLayers(layers).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLibraryWithLinksForTests creates a new library with links for tests
func NewLibraryWithLinksForTests(layers layers.Layers, links links.Links) Library {
	ins, err := NewBuilder().Create().WithLayers(layers).WithLinks(links).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
