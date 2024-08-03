package profiles

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

// Profile represents a profile
type Profile interface {
	Hash() hash.Hash
	Handle() string
	Description() string
}
