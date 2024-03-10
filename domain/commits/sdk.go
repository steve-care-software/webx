package commits

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/commits/contents"
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents the commit builder
type Builder interface {
	Create() Builder
	WithContent(content contents.Content) Builder
	WithSignature(signature signers.Signature) Builder
	WithIndex(index uint) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Content() contents.Content
	Signature() signers.Signature
	Index() uint
}
