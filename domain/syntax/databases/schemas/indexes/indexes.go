package indexes

import "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"

type indexes struct {
	hash hash.Hash
	list []Index
}

func createIndexes(
	hash hash.Hash,
	list []Index,
) Indexes {
	out := indexes{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *indexes) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *indexes) List() []Index {
	return obj.list
}
