package transfers

import (
	"github.com/steve-care-software/webx/engine/domain/hash"
)

// Transfer represents a token transfer
type Transfer interface {
	Hash() hash.Hash
	Token() hash.Hash
	NewOwner() []hash.Hash
}
