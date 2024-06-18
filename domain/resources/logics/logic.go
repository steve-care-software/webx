package logics

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/resources/logics/bridges"
	"github.com/steve-care-software/datastencil/domain/resources/logics/references"
)

type logic struct {
	hash       hash.Hash
	link       links.Link
	bridges    bridges.Bridges
	references references.References
}

func createLogic(
	hash hash.Hash,
	link links.Link,
	bridges bridges.Bridges,
) Logic {
	return createLogicInternally(hash, link, bridges, nil)
}

func createLogicWithReferences(
	hash hash.Hash,
	link links.Link,
	bridges bridges.Bridges,
	references references.References,
) Logic {
	return createLogicInternally(hash, link, bridges, references)
}

func createLogicInternally(
	hash hash.Hash,
	link links.Link,
	bridges bridges.Bridges,
	references references.References,
) Logic {
	out := logic{
		hash:       hash,
		link:       link,
		bridges:    bridges,
		references: references,
	}

	return &out
}

// Hash returns the hash
func (obj *logic) Hash() hash.Hash {
	return obj.hash
}

// Link returns the link
func (obj *logic) Link() links.Link {
	return obj.link
}

// Bridges returns the layers
func (obj *logic) Bridges() bridges.Bridges {
	return obj.bridges
}

// HasReferences returns true if there is references, false otherwise
func (obj *logic) HasReferences() bool {
	return obj.references != nil
}

// References returns the references, if any
func (obj *logic) References() references.References {
	return obj.references
}
