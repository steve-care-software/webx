package relationals

import "github.com/steve-care-software/datastencil/domain/hash"

type relational struct {
	hash  hash.Hash
	isAnd bool
	isOr  bool
}

func createRelationalWithAnd(
	hash hash.Hash,
) Relational {
	return createRelationalInternally(hash, true, false)
}

func createRelationalWithOr(
	hash hash.Hash,
) Relational {
	return createRelationalInternally(hash, false, true)
}

func createRelationalInternally(
	hash hash.Hash,
	isAnd bool,
	isOr bool,
) Relational {
	out := relational{
		hash:  hash,
		isAnd: isAnd,
		isOr:  isOr,
	}

	return &out
}

// Hash returns the hash
func (obj *relational) Hash() hash.Hash {
	return nil
}

// IsAnd returns true if and, false otherwise
func (obj *relational) IsAnd() bool {
	return obj.isAnd
}

// IsOr returns true if or, false otherwise
func (obj *relational) IsOr() bool {
	return obj.isOr
}
