package grammars

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

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

// Points returns the amount of points a line contains
func (obj *line) Points() uint {
	amount := uint(0)
	for _, oneElement := range obj.elements {
		amount += oneElement.Points()
	}

	return amount
}

// Elements returns the elements
func (obj *line) Elements() []Element {
	return obj.elements
}
