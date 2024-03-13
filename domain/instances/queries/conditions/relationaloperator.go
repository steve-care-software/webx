package conditions

import "github.com/steve-care-software/datastencil/domain/hash"

type relationalOperator struct {
	hash  hash.Hash
	isAnd bool
	isOr  bool
}

func createRelationalOperatorWithAnd(
	hash hash.Hash,
) RelationalOperator {
	return createRelationalOperatorInternally(hash, true, false)
}

func createRelationalOperatorWithOr(
	hash hash.Hash,
) RelationalOperator {
	return createRelationalOperatorInternally(hash, false, true)
}

func createRelationalOperatorInternally(
	hash hash.Hash,
	isAnd bool,
	isOr bool,
) RelationalOperator {
	out := relationalOperator{
		hash:  hash,
		isAnd: isAnd,
		isOr:  isOr,
	}

	return &out
}

// Hash returns the hash
func (obj *relationalOperator) Hash() hash.Hash {
	return nil
}

// IsAnd returns true if and, false otherwise
func (obj *relationalOperator) IsAnd() bool {
	return obj.isAnd
}

// IsOr returns true if or, false otherwise
func (obj *relationalOperator) IsOr() bool {
	return obj.isOr
}
