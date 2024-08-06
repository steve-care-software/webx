package tokens

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

type token struct {
	hash        hash.Hash
	elements    elements.Elements
	cardinality cardinalities.Cardinality
	omissiom    omissions.Omission
}

func createToken(
	hash hash.Hash,
	elements elements.Elements,
	cardinality cardinalities.Cardinality,
) Token {
	return createTokenInternally(hash, elements, cardinality, nil)
}

func createTokenWithOmission(
	hash hash.Hash,
	elements elements.Elements,
	cardinality cardinalities.Cardinality,
	omissiom omissions.Omission,
) Token {
	return createTokenInternally(hash, elements, cardinality, omissiom)
}

func createTokenInternally(
	hash hash.Hash,
	elements elements.Elements,
	cardinality cardinalities.Cardinality,
	omissiom omissions.Omission,
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
func (obj *token) Elements() elements.Elements {
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
func (obj *token) Omission() omissions.Omission {
	return obj.omissiom
}
