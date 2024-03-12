package commits

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents the commit builder
type Builder interface {
	Create() Builder
	WithContent(content contents.Content) Builder
	WithSignature(signature signers.Signature) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Content() contents.Content
	Signature() signers.Signature
	Index() uint
	PublicKey() (signers.PublicKey, error)
}
