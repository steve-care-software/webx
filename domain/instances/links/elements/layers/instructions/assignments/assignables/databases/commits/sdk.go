package commits

import "github.com/steve-care-software/datastencil/domain/hash"

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Description() string
	Actions() []string
	HashParent() bool
	Parent() string
}
