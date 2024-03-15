package commits

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
)

type content struct {
	hash     hash.Hash
	actions  actions.Actions
	previous Commit
}

func createContent(
	hash hash.Hash,
	actions actions.Actions,
) Content {
	return createContentInternally(hash, actions, nil)
}

func createContentWithPrevious(
	hash hash.Hash,
	actions actions.Actions,
	previous Commit,
) Content {
	return createContentInternally(hash, actions, previous)
}

func createContentInternally(
	hash hash.Hash,
	actions actions.Actions,
	previous Commit,
) Content {
	out := content{
		hash:     hash,
		actions:  actions,
		previous: previous,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Actions returns the actions
func (obj *content) Actions() actions.Actions {
	return obj.actions
}

// HasPrevious returns true if there is a previous, false otherwise
func (obj *content) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous
func (obj *content) Previous() Commit {
	return obj.previous
}
