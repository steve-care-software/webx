package operators

import "github.com/steve-care-software/datastencil/domain/hash"

type operator struct {
	hash  hash.Hash
	isAnd bool
	isOr  bool
	isXor bool
}

func createOperatorWithIsAnd(
	hash hash.Hash,
) Operator {
	return createOperatorInternally(hash, true, false, false)
}

func createOperatorWithIsOr(
	hash hash.Hash,
) Operator {
	return createOperatorInternally(hash, false, true, false)
}

func createOperatorWithIsXor(
	hash hash.Hash,
) Operator {
	return createOperatorInternally(hash, false, false, true)
}

func createOperatorInternally(
	hash hash.Hash,
	isAnd bool,
	isOr bool,
	isXor bool,
) Operator {
	out := operator{
		hash:  hash,
		isAnd: isAnd,
		isOr:  isOr,
		isXor: isXor,
	}

	return &out
}

// Hash returns the hash
func (obj *operator) Hash() hash.Hash {
	return obj.hash
}

// IsAnd returns true if there is an and, false otherwise
func (obj *operator) IsAnd() bool {
	return obj.isAnd
}

// IsOr returns true if there is an or, false otherwise
func (obj *operator) IsOr() bool {
	return obj.isOr
}

// IsXor returns true if there is a xor, false otherwise
func (obj *operator) IsXor() bool {
	return obj.isXor
}
