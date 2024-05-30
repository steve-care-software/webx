package logics

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
)

type logic struct {
	hash    hash.Hash
	link    links.Link
	bridges bridges.Bridges
}

func createLogic(
	hash hash.Hash,
	link links.Link,
	bridges bridges.Bridges,
) Logic {
	out := logic{
		hash:    hash,
		link:    link,
		bridges: bridges,
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
