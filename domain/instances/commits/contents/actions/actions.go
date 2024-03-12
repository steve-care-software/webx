package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type actions struct {
	hash hash.Hash
	mp   map[string]Action
	list []Action
}

func createActions(
	hash hash.Hash,
	mp map[string]Action,
	list []Action,
) Actions {
	out := actions{
		hash: hash,
		mp:   mp,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *actions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *actions) List() []Action {
	return obj.list
}
