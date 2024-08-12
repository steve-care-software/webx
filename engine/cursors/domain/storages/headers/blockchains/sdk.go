package blockchains

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/blockchains/roots"
)

// Blockchain represents a blockchain
type Blockchain interface {
	Root() roots.Root
	HasHead() bool
	Head() delimiters.Delimiter
}
