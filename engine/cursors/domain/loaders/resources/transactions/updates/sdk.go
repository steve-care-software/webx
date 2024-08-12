package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// Update represents an update
type Update interface {
	Hash() hash.Hash
	DelimiterIndex() uint64
	HasBlacklist() bool
	Blacklist() []hash.Hash
	HasWhitelist() bool
	Whitelist() []hash.Hash
}
