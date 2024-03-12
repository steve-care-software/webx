package previous

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions"
)

type previous struct {
	hash hash.Hash
	root actions.Actions
	prev Previous
}

func createPreviousWithRoot(
	hash hash.Hash,
	root actions.Actions,
) Previous {
	return createPreviousInternally(hash, root, nil)
}

func createPreviousWithPrevious(
	hash hash.Hash,
	prev Previous,
) Previous {
	return createPreviousInternally(hash, nil, prev)
}

func createPreviousInternally(
	hash hash.Hash,
	root actions.Actions,
	prev Previous,
) Previous {
	out := previous{
		hash: hash,
		root: root,
		prev: prev,
	}

	return &out
}

// Hash returns the hash
func (obj *previous) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *previous) Index() uint {
	if obj.IsRoot() {
		return 0
	}

	return obj.prev.Index() + 1
}

// IsRoot returns true if there is root, false otherwise
func (obj *previous) IsRoot() bool {
	return obj.root != nil
}

// Root returns the root, if any
func (obj *previous) Root() actions.Actions {
	return obj.root
}

// IsPrevious returns true if there is a previous, false otherwise
func (obj *previous) IsPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous, if any
func (obj *previous) Previous() Previous {
	return obj.prev
}
