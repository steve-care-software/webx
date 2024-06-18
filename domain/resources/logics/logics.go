package logics

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type logics struct {
	hash hash.Hash
	list []Logic
}

func createLogics(
	hash hash.Hash,
	list []Logic,
) Logics {
	out := logics{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *logics) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *logics) List() []Logic {
	return obj.list
}
