package references

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type references struct {
	hash hash.Hash
	mp   map[string]Reference
	list []Reference
}

func createReferences(
	hash hash.Hash,
	mp map[string]Reference,
	list []Reference,
) References {
	out := references{
		hash: hash,
		mp:   mp,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *references) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *references) List() []Reference {
	return obj.list
}

// Fetch fetches a reference by hash
func (obj *references) Fetch(name string) (Reference, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the requested Reference (variable: %s) does not exists", name)
	return nil, errors.New(str)
}
