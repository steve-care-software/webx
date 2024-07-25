package cardinalities

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type cardinality struct {
	hash hash.Hash
	min  uint
	pMax *uint
}

func createCardinality(
	hash hash.Hash,
	min uint,
) Cardinality {
	return createCardinalityInternally(hash, min, nil)
}

func createCardinalityWithMax(
	hash hash.Hash,
	min uint,
	pMax *uint,
) Cardinality {
	return createCardinalityInternally(hash, min, pMax)
}

func createCardinalityInternally(
	hash hash.Hash,
	min uint,
	pMax *uint,
) Cardinality {
	out := cardinality{
		hash: hash,
		min:  min,
		pMax: pMax,
	}

	return &out
}

// Hash returns the hash
func (obj *cardinality) Hash() hash.Hash {
	return obj.hash
}

// Min returns the min
func (obj *cardinality) Min() uint {
	return obj.min
}

// HasMax returns true if there is a max, false otherwise
func (obj *cardinality) HasMax() bool {
	return obj.pMax != nil
}

// Max returns the max, if any
func (obj *cardinality) Max() *uint {
	return obj.pMax
}
