package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
)

type link struct {
	hash   hash.Hash
	input  []byte
	source links.Link
	layers layers.Layers
	next   Link
}

func createLink(
	hash hash.Hash,
	input []byte,
	source links.Link,
) Link {
	return createLinkIntrnally(hash, input, source, nil, nil)
}

func createLinkWithLayers(
	hash hash.Hash,
	input []byte,
	source links.Link,
	layers layers.Layers,
) Link {
	return createLinkIntrnally(hash, input, source, layers, nil)
}

func createLinkWithNext(
	hash hash.Hash,
	input []byte,
	source links.Link,
	next Link,
) Link {
	return createLinkIntrnally(hash, input, source, nil, next)
}

func createLinkWithLayersAndNext(
	hash hash.Hash,
	input []byte,
	source links.Link,
	layers layers.Layers,
	next Link,
) Link {
	return createLinkIntrnally(hash, input, source, layers, next)
}

func createLinkIntrnally(
	hash hash.Hash,
	input []byte,
	source links.Link,
	layers layers.Layers,
	next Link,
) Link {
	out := link{
		hash:   hash,
		input:  input,
		source: source,
		layers: layers,
		next:   next,
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

// HasNext returns true if there is a next link, false otherwise
func (obj *link) HasNext() bool {
	return obj.next != nil
}

// Next returns the next link, if any
func (obj *link) Next() Link {
	return obj.next
}
