package integers

import "github.com/steve-care-software/datastencil/domain/hash"

type integer struct {
	hash          hash.Hash
	isSmallerThan bool
	isBiggerThan  bool
	isEqual       bool
}

func createIntegerWithSmallerThan(
	hash hash.Hash,
) Integer {
	return createIntegerInternally(
		hash,
		true,
		false,
		false,
	)
}

func createIntegerWithSmallerThanAndEqual(
	hash hash.Hash,
) Integer {
	return createIntegerInternally(
		hash,
		true,
		false,
		true,
	)
}

func createIntegerWithBiggerThan(
	hash hash.Hash,
) Integer {
	return createIntegerInternally(
		hash,
		false,
		true,
		false,
	)
}

func createIntegerWithBiggerThanAndEqual(
	hash hash.Hash,
) Integer {
	return createIntegerInternally(
		hash,
		false,
		true,
		true,
	)
}

func createIntegerWithEqual(
	hash hash.Hash,
) Integer {
	return createIntegerInternally(
		hash,
		false,
		false,
		true,
	)
}

func createIntegerInternally(
	hash hash.Hash,
	isSmallerThan bool,
	isBiggerThan bool,
	isEqual bool,
) Integer {
	out := integer{
		hash:          hash,
		isSmallerThan: isSmallerThan,
		isBiggerThan:  isBiggerThan,
		isEqual:       isEqual,
	}

	return &out
}

// Hash returns the hash
func (obj *integer) Hash() hash.Hash {
	return obj.hash
}

// IsSmallerThan returns true if smaller than, false otherwise
func (obj *integer) IsSmallerThan() bool {
	return obj.isSmallerThan
}

// IsBiggerThan returns true if bigger than, false otherwise
func (obj *integer) IsBiggerThan() bool {
	return obj.isBiggerThan
}

// IsEqual returns true if equal, false otherwise
func (obj *integer) IsEqual() bool {
	return obj.isEqual
}
