package blocks

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Transactions() delimiters.Delimiter
	Answer() []byte
	Result() hash.Hash
}
