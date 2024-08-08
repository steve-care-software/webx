package locks

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/actions/transfers/contents/origins"
)

// Lock represents a lock
type Lock interface {
	Hash() hash.Hash
	Origin() origins.Origin
	Secret() hash.Hash
}
