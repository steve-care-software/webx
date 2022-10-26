package criterias

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type delimiter struct {
	hash    hash.Hash
	index   uint
	pAmount *uint
}

func createDelimiter(
	hash hash.Hash,
	index uint,
) Delimiter {
	return createDelimiterInternally(hash, index, nil)
}

func createDelimiterWithAmount(
	hash hash.Hash,
	index uint,
	pAmount *uint,
) Delimiter {
	return createDelimiterInternally(hash, index, pAmount)
}

func createDelimiterInternally(
	hash hash.Hash,
	index uint,
	pAmount *uint,
) Delimiter {
	out := delimiter{
		hash:    hash,
		index:   index,
		pAmount: pAmount,
	}

	return &out
}

// Hash returns the hash
func (obj *delimiter) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *delimiter) Index() uint {
	return obj.index
}

// HasAmount returns true if there is an amount, false otherwise
func (obj *delimiter) HasAmount() bool {
	return obj.pAmount != nil
}

// Amount returns the amount, if any
func (obj *delimiter) Amount() *uint {
	return obj.pAmount
}
