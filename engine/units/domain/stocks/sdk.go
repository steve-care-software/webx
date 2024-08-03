package stocks

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/identities/profiles"
)

// Stocks represents stocks
type Stocks interface {
	Hash() hash.Hash
	List() []Stock
}

// Stock represents a stock
type Stock interface {
	Hash() hash.Hash
	Profile() profiles.Profile
	Amount() uint64
}
