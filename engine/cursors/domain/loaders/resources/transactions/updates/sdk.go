package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
)

// Update represents an update
type Update interface {
	Hash() hash.Hash
	Content() Content
	Vote() signers.Vote
}

// Content represents an update content
type Content interface {
	Hash() hash.Hash
	Name() string
	DelimiterIndex() uint64
	HasBlacklistAddition() bool
	BlacklistAddition() []hash.Hash
	HasBlacklistRemoval() bool
	BlacklistRemoval() []hash.Hash
	HasWhitelistAddition() bool
	WhitelistAddition() []hash.Hash
	HasWhitelistRemoval() bool
	WhitelistRemoval() []hash.Hash
}
