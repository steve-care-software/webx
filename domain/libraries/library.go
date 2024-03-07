package libraries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
)

type library struct {
	hash   hash.Hash
	layers layers.Layers
	links  links.Links
}

func createLibrary(
	hash hash.Hash,
	layers layers.Layers,
) Library {
	return createLibraryInternally(hash, layers, nil)
}

func createLibraryWithLinks(
	hash hash.Hash,
	layers layers.Layers,
	links links.Links,
) Library {
	return createLibraryInternally(hash, layers, links)
}

func createLibraryInternally(
	hash hash.Hash,
	layers layers.Layers,
	links links.Links,
) Library {
	out := library{
		hash:   hash,
		layers: layers,
		links:  links,
	}

	return &out
}

// Hash returns the hash
func (obj *library) Hash() hash.Hash {
	return obj.hash
}

// Layers returns the layers, if any
func (obj *library) Layers() layers.Layers {
	return obj.layers
}

// HasLinks returns true if there is links, false otherwise
func (obj *library) HasLinks() bool {
	return obj.links != nil
}

// Links returns the links, if any
func (obj *library) Links() links.Links {
	return obj.links
}
