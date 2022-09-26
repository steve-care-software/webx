package units

import "github.com/steve-care-software/syntax/domain/identity/cryptography/hash"

type units struct {
	hash hash.Hash
	list []Unit
}

func createUnits(
	hash hash.Hash,
	list []Unit,
) Units {
	out := units{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *units) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *units) Amount() uint64 {
	amount := uint64(0)
	for _, oneUnit := range obj.list {
		amount += oneUnit.Content().Amount()
	}

	return amount
}

// List returns the list
func (obj *units) List() []Unit {
	return obj.list
}
