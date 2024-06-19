package queries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type queries struct {
	hash hash.Hash
	list []Query
}

func createQueries(
	hash hash.Hash,
	list []Query,
) Queries {
	out := queries{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *queries) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *queries) List() []Query {
	return obj.list
}
