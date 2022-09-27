package units

import "github.com/steve-care-software/syntax/domain/identity/cryptography/hash"

type content struct {
	hash     hash.Hash
	amount   uint64
	owner    []hash.Hash
	previous Previous
}

func createContent(
	hash hash.Hash,
	amount uint64,
	owner []hash.Hash,
	previous Previous,
) Content {
	out := content{
		hash:     hash,
		amount:   amount,
		owner:    owner,
		previous: previous,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *content) Amount() uint64 {
	return obj.amount
}

// Owner returns the owner
func (obj *content) Owner() []hash.Hash {
	return obj.owner
}

// Previous returns the previous
func (obj *content) Previous() Previous {
	return obj.previous
}
