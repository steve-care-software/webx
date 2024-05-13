package fetches

import "github.com/steve-care-software/datastencil/domain/hash"

type fetch struct {
	hash  hash.Hash
	list  string
	index string
}

func createFetch(
	hash hash.Hash,
	list string,
	index string,
) Fetch {
	out := fetch{
		hash:  hash,
		list:  list,
		index: index,
	}

	return &out
}

// Hash returns the hash
func (obj *fetch) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *fetch) List() string {
	return obj.list
}

// Index returns the index
func (obj *fetch) Index() string {
	return obj.index
}
