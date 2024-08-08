package origins

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Origin represents a transfer origin
type Origin interface {
	Hash() hash.Hash
	IsPurse() bool
	Purse() hash.Hash
	IsWallet() bool
	Wallet() hash.Hash
}
