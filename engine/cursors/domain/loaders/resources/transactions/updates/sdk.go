package updates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSigner(signer signers.Signer) Builder
	WithRing(ring []signers.PublicKey) Builder
	WithData(data []byte) Builder
	WithWhiteListAddition(addition []hash.Hash) Builder
	WithWhiteListRemoval(removal []hash.Hash) Builder
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
