package conditions

import "github.com/steve-care-software/datastencil/domain/hash"

type operator struct {
	hash       hash.Hash
	isEqual    bool
	relational RelationalOperator
	integer    IntegerOperator
}

func createOperatorWithEqual(
	hash hash.Hash,
) Operator {
	return createOperatorInternally(hash, true, nil, nil)
}

func createOperatorWithRelational(
	hash hash.Hash,
	relational RelationalOperator,
) Operator {
	return createOperatorInternally(hash, false, relational, nil)
}

func createOperatorWithInteger(
	hash hash.Hash,
	integer IntegerOperator,
) Operator {
	return createOperatorInternally(hash, false, nil, integer)
}

func createOperatorInternally(
	hash hash.Hash,
	isEqual bool,
	relational RelationalOperator,
	integer IntegerOperator,
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
func (obj *operator) Relational() RelationalOperator {
	return obj.relational
}

// IsInteger returns true if integer, false otherwise
func (obj *operator) IsInteger() bool {
	return obj.integer != nil
}

// Integer returns the integer operaotr, if any
func (obj *operator) Integer() IntegerOperator {
	return obj.integer
}
