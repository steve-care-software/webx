package genesis

import "github.com/steve-care-software/syntax/domain/identity/cryptography/hash"

type genesis struct {
	hash        hash.Hash
	ticker      string
	description string
	supply      uint64
	owner       []hash.Hash
}

func createGenesis(
	hash hash.Hash,
	ticker string,
	description string,
	supply uint64,
	owner []hash.Hash,
) Genesis {
	out := genesis{
		hash:        hash,
		ticker:      ticker,
		description: description,
		supply:      supply,
		owner:       owner,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Ticker returns the ticker
func (obj *genesis) Ticker() string {
	return obj.ticker
}

// Description returns the description
func (obj *genesis) Description() string {
	return obj.description
}

// Supply returns the supply
func (obj *genesis) Supply() uint64 {
	return obj.supply
}

// Owner returns the owner
func (obj *genesis) Owner() []hash.Hash {
	return obj.owner
}
