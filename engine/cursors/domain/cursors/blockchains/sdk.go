package blockchains

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/blockchains/bags"
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/blockchains/blocks"
)

// Blockchain represents a blockchain
type Blockchain interface {
	HasBag() bool
	Bag() bags.Bag
	HasBlock() bool
	Block() blocks.Block
}
