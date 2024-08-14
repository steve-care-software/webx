package transfers

import (
	"github.com/steve-care-software/webx/engine/domain/blockchains/hash"
)

// Transfer represents a token transfer
type Transfer interface {
	Hash() hash.Hash
	Token() hash.Hash
	NewOwner() []hash.Hash
}
