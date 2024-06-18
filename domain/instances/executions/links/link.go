package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links"
)

type link struct {
	hash   hash.Hash
	input  []byte
	source links.Link
	layers layers.Layers
}

func createLink(
	hash hash.Hash,
	input []byte,
	source links.Link,
) Link {
	return createLinkIntrnally(hash, input, source, nil)
}

func createLinkWithLayers(
	hash hash.Hash,
	input []byte,
	source links.Link,
	layers layers.Layers,
) Link {
	return createLinkIntrnally(hash, input, source, layers)
}

func createLinkIntrnally(
	hash hash.Hash,
	input []byte,
	source links.Link,
	layers layers.Layers,
) Link {
	out := link{
		hash:   hash,
		input:  input,
		source: source,
		layers: layers,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *link) Input() []byte {
	return obj.input
}

// Source returns the source
func (obj *link) Source() links.Link {
	return obj.source
}

// HasLayers returns true if there is layers, false otherwise
func (obj *link) HasLayers() bool {
	return obj.layers != nil
}

// Layers returns layers, if any
func (obj *link) Layers() layers.Layers {
	return obj.layers
}
