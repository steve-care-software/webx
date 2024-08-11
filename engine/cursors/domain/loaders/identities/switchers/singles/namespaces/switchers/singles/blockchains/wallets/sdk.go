package wallets

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Wallets represents wallets
type Wallets interface {
	Hash() hash.Hash
	List() []Wallet
}

// Wallet represents a wallet
type Wallet interface {
	Hash() hash.Hash
	HasIncoming() bool
	//Incoming() transfers.Transfer
	HasOutgoing() bool
	Outgoing() bool
	//Transfer() transfers.Transfer
}
