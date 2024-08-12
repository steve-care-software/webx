package transactions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
}
