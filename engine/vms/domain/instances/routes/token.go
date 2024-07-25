package routes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/cardinalities"
)

type token struct {
	hash        hash.Hash
	elements    Elements
	cardinality cardinalities.Cardinality
	omissiom    Omission
}

func createToken(
	hash hash.Hash,
	elements Elements,
	cardinality cardinalities.Cardinality,
) Token {
	return createTokenInternally(hash, elements, cardinality, nil)
}

func createTokenWithOmission(
	hash hash.Hash,
	elements Elements,
	cardinality cardinalities.Cardinality,
	omissiom Omission,
) Token {
	return createTokenInternally(hash, elements, cardinality, omissiom)
}

func createTokenInternally(
	hash hash.Hash,
	elements Elements,
	cardinality cardinalities.Cardinality,
	omissiom Omission,
) Token {
	out := token{
		hash:        hash,
		elements:    elements,
		cardinality: cardinality,
		omissiom:    omissiom,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Elements returns the elements
func (obj *token) Elements() Elements {
	return obj.elements
}

// Cardinality returns the cardinality
func (obj *token) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}

// HasOmission returns true if there is an omission, false otherwise
func (obj *token) HasOmission() bool {
	return obj.omissiom != nil
}

// Omission returns the omission, if any
func (obj *token) Omission() Omission {
	return obj.omissiom
}
