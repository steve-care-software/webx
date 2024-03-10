package conditions

import "github.com/steve-care-software/datastencil/domain/hash"

type integerOperator struct {
	hash          hash.Hash
	isSmallerThan bool
	isBiggerThan  bool
	isEqual       bool
}

func createIntegerOperatorWithSmallerThan(
	hash hash.Hash,
) IntegerOperator {
	return createIntegerOperatorInternally(
		hash,
		true,
		false,
		false,
	)
}

func createIntegerOperatorWithSmallerThanAndEqual(
	hash hash.Hash,
) IntegerOperator {
	return createIntegerOperatorInternally(
		hash,
		true,
		false,
		true,
	)
}

func createIntegerOperatorWithBiggerThan(
	hash hash.Hash,
) IntegerOperator {
	return createIntegerOperatorInternally(
		hash,
		false,
		true,
		false,
	)
}

func createIntegerOperatorWithBiggerThanAndEqual(
	hash hash.Hash,
) IntegerOperator {
	return createIntegerOperatorInternally(
		hash,
		false,
		true,
		true,
	)
}

func createIntegerOperatorWithEqual(
	hash hash.Hash,
) IntegerOperator {
	return createIntegerOperatorInternally(
		hash,
		false,
		false,
		true,
	)
}

func createIntegerOperatorInternally(
	hash hash.Hash,
	isSmallerThan bool,
	isBiggerThan bool,
	isEqual bool,
) IntegerOperator {
	out := integerOperator{
		hash:          hash,
		isSmallerThan: isSmallerThan,
		isBiggerThan:  isBiggerThan,
		isEqual:       isEqual,
	}

	return &out
}

// Hash returns the hash
func (obj *integerOperator) Hash() hash.Hash {
	return obj.hash
}

// IsSmallerThan returns true if smaller than, false otherwise
func (obj *integerOperator) IsSmallerThan() bool {
	return obj.isSmallerThan
}

// IsBiggerThan returns true if bigger than, false otherwise
func (obj *integerOperator) IsBiggerThan() bool {
	return obj.isBiggerThan
}

// IsEqual returns true if equal, false otherwise
func (obj *integerOperator) IsEqual() bool {
	return obj.isEqual
}
