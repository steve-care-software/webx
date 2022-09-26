package genesis

import "github.com/steve-care-software/syntax/domain/identity/cryptography/hash"

type genesis struct {
	hash   hash.Hash
	supply uint64
	owner  []hash.Hash
}

func createGenesis(
	hash hash.Hash,
	supply uint64,
	owner []hash.Hash,
) Genesis {
	out := genesis{
		hash:   hash,
		supply: supply,
		owner:  owner,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Supply returns the supply
func (obj *genesis) Supply() uint64 {
	return obj.supply
}

// Owner returns the owner
func (obj *genesis) Owner() []hash.Hash {
	return obj.owner
}
