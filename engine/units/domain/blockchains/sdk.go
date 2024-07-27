package blockchains

import (
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/genesis"
)

// Blockchain represents a blockchain
type Blockchain interface {
	Identifier() string
	Genesis() genesis.Genesis
	HasPeers() bool
	Peers() []string
	HasHead() bool
	Head() blocks.Block
}
