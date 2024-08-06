package identities

import "github.com/steve-care-software/webx/engine/cursors/domain/cursors/identities/wallets"

// Identity represents an identity
type Identity interface {
	Name() string
	HasWallet() bool
	Wallet() wallets.Wallet
}
