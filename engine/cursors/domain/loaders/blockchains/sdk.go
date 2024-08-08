package blockchains

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/wallets"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
)

// Blockchain represents the blockchain
type Blockchain interface {
	Root() storages.Storage
	HasHad() bool
	Head() storages.Storage
	HasWallets() bool
	Wallets() wallets.Wallets
}
