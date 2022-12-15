package tokens

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

type line struct {
	elements []hash.Hash
}

func createLine(
	elements []hash.Hash,
) Line {
	out := line{
		elements: elements,
	}

	return &out
}

// Elements returns the elements hashes
func (obj *line) Elements() []hash.Hash {
	return obj.elements
}
