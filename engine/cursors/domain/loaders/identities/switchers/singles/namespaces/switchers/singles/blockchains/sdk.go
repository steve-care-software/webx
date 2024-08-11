package blockchains

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/blocks"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/roots"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/blockchains/wallets"
)

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithRoot(root roots.Root) Builder
	WithHead(head blocks.Block) Builder
	WithWallets(wallets wallets.Wallets) Builder
	Now() (Blockchain, error)
}

// Blockchain represents the blockchain
type Blockchain interface {
	Root() roots.Root
	HasHead() bool
	Head() blocks.Block
	HasWallets() bool
	Wallets() wallets.Wallets
}
