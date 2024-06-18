package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/references"
)

type link struct {
	hash       hash.Hash
	elements   elements.Elements
	references references.References
}

func createLink(
	hash hash.Hash,
	elements elements.Elements,
) Link {
	return createLinkInternally(hash, elements, nil)
}

func createLinkWithReferences(
	hash hash.Hash,
	elements elements.Elements,
	references references.References,
) Link {
	return createLinkInternally(hash, elements, references)
}

func createLinkInternally(
	hash hash.Hash,
	elements elements.Elements,
	references references.References,
) Link {
	out := link{
		hash:       hash,
		elements:   elements,
		references: references,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Elements returns the elements
func (obj *link) Elements() elements.Elements {
	return obj.elements
}

// HasReferences returns true if there is references, false otherwise
func (obj *link) HasReferences() bool {
	return obj.references != nil
}

// References returns the references, if any
func (obj *link) References() references.References {
	return obj.references
}
