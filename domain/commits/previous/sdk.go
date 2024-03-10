package previous

import (
	"github.com/steve-care-software/datastencil/domain/commits/actions"
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Previous represents a previous
type Previous interface {
	Hash() hash.Hash
	IsRoot() bool
	Root() actions.Actions
	IsPrevious() bool
	Previous() Previous
}
