package libraries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/references"
)

type library struct {
	hash       hash.Hash
	layers     layers.Layers
	links      links.Links
	references references.References
}

func createLibrary(
	hash hash.Hash,
	layers layers.Layers,
) Library {
	return createLibraryInternally(hash, layers, nil, nil)
}

func createLibraryWithLinks(
	hash hash.Hash,
	layers layers.Layers,
	links links.Links,
) Library {
	return createLibraryInternally(hash, layers, links, nil)
}

func createLibraryWithReferences(
	hash hash.Hash,
	layers layers.Layers,
	references references.References,
) Library {
	return createLibraryInternally(hash, layers, nil, references)
}

func createLibraryWithLinksAndReferences(
	hash hash.Hash,
	layers layers.Layers,
	links links.Links,
	references references.References,
) Library {
	return createLibraryInternally(hash, layers, links, references)
}

func createLibraryInternally(
	hash hash.Hash,
	layers layers.Layers,
	links links.Links,
	references references.References,
) Library {
	out := library{
		hash:       hash,
		layers:     layers,
		links:      links,
		references: references,
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

// HasReferences returns true if there is references, false otherwise
func (obj *library) HasReferences() bool {
	return obj.references != nil
}

// References returns the references, if any
func (obj *library) References() references.References {
	return obj.references
}
