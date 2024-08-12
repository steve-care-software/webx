package inserts

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Insert represents an insert
type Insert interface {
	Hash() hash.Hash
	Name() string
	Bytes() []byte
	Whitelist() []hash.Hash
}
