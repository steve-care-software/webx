package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
)

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSigner(signer signers.Signer) Builder
	WithData(data []byte) Builder
	WithWhiteListAddition(wlAddition []hash.Hash) Builder
	WithWhiteListRemoval(wlRemoval []hash.Hash) Builder
	Now() (Update, error)
}

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
	HasData() bool
	Data() []byte
	HasWhitelistAddition() bool
	WhitelistAddition() []hash.Hash
	HasWhitelistRemoval() bool
	WhitelistRemoval() []hash.Hash
}