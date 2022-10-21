package grammars

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type line struct {
	hash     hash.Hash
	elements []Element
}

func createLine(
	hash hash.Hash,
	elements []Element,
) Line {
	out := line{
		hash:     hash,
		elements: elements,
	}

	return &out
}

// Hash returns the hash
func (obj *line) Hash() hash.Hash {
	return obj.hash
}

// Elements returns the elements
func (obj *line) Elements() []Element {
	return obj.elements
}
