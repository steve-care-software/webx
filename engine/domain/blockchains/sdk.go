package blockchains

import (
	"github.com/steve-care-software/webx/engine/domain/blockchains/hash"
	"github.com/steve-care-software/webx/engine/domain/blockchains/roots"
)

// Blockchain represents a blockchain
type Blockchain interface {
	Root() roots.Root
	Head() hash.Hash
	HasPeers() bool
	Peers() []string
}
