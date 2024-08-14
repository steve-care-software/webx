package contents

import (
	"time"

	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions/contents/containers"
	"github.com/steve-care-software/webx/engine/domain/blockchains/hash"
)

// Content represents the content of a transaction
type Content interface {
	Hash() hash.Hash
	Containers() containers.Containers
	Condition() hash.Hash
	ExpireIn() uint
	CreatedOn() time.Time
}
