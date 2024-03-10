package commits

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/commits/actions"
	"github.com/steve-care-software/datastencil/domain/commits/previous"
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Content() Content
	Signature() signers.Signature
	Index() uint
}

// Content represents a commit content
type Content interface {
	Hash() hash.Hash
	Actions() actions.Actions
	Previous() previous.Previous
}
