package operators

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/integers"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/relationals"
)

type operator struct {
	hash       hash.Hash
	isEqual    bool
	relational relationals.Relational
	integer    integers.Integer
}

func createOperatorWithEqual(
	hash hash.Hash,
) Operator {
	return createOperatorInternally(hash, true, nil, nil)
}

func createOperatorWithRelational(
	hash hash.Hash,
	relational relationals.Relational,
) Operator {
	return createOperatorInternally(hash, false, relational, nil)
}

func createOperatorWithInteger(
	hash hash.Hash,
	integer integers.Integer,
) Operator {
	return createOperatorInternally(hash, false, nil, integer)
}

func createOperatorInternally(
	hash hash.Hash,
	isEqual bool,
	relational relationals.Relational,
	integer integers.Integer,
) Operator {
	out := operator{
		hash:       hash,
		isEqual:    isEqual,
		relational: relational,
		integer:    integer,
	}

	return &out
}

// Hash returns the hash
func (obj *operator) Hash() hash.Hash {
	return obj.hash
}

// IsEqual returns true if equal, false otherwise
func (obj *operator) IsEqual() bool {
	return obj.isEqual
}

// IsRelational returns true if relational, false otherwise
func (obj *operator) IsRelational() bool {
	return obj.relational != nil
}

// Relational returns the relational operaotr, if any
func (obj *operator) Relational() relationals.Relational {
	return obj.relational
}

// IsInteger returns true if integer, false otherwise
func (obj *operator) IsInteger() bool {
	return obj.integer != nil
}

// Integer returns the integer operaotr, if any
func (obj *operator) Integer() integers.Integer {
	return obj.integer
}
