package inserts

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Insert represents an insert
type Insert interface {
	Hash() hash.Hash
	Bytes() []byte
	HasBlacklist() bool
	Blacklist() []hash.Hash
	HasWhitelist() bool
	Whitelist() []hash.Hash
}
