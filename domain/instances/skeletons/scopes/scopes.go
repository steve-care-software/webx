package scopes

import "github.com/steve-care-software/datastencil/domain/hash"

type scopes struct {
	hash hash.Hash
	list []Scope
}

func createScopes(
	hash hash.Hash,
	list []Scope,
) Scopes {
	out := scopes{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *scopes) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *scopes) List() []Scope {
	return obj.list
}

// Contains returns true if the path is contained in the scope, false otherwise
func (obj *scopes) Contains(path []string) bool {
	for _, oneScope := range obj.list {
		if oneScope.Contains(path) {
			return true
		}
	}

	return false
}
