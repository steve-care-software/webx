package logics

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
)

type logic struct {
	hash   hash.Hash
	link   links.Link
	layers layers.Layers
}

func createLogic(
	hash hash.Hash,
	link links.Link,
	layers layers.Layers,
) Logic {
	out := logic{
		hash:   hash,
		link:   link,
		layers: layers,
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

// Layers returns the layers
func (obj *logic) Layers() layers.Layers {
	return obj.layers
}
