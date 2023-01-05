package grammars

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type line struct {
	hash       hash.Hash
	containers []Container
}

func createLine(
	hash hash.Hash,
	containers []Container,
) Line {
	out := line{
		hash:       hash,
		containers: containers,
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
	for _, oneContainer := range obj.containers {
		amount += oneContainer.Points()
	}

	return amount
}

// Containers returns the containers
func (obj *line) Containers() []Container {
	return obj.containers
}
