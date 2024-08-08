package contents

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/locks"
)

// Content represents a transfer content
type Content interface {
	Hash() hash.Hash
	Lock() locks.Lock
	Condition() hash.Hash
	ConditionExpireIn() uint64
}
