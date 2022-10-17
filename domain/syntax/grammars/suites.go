package grammars

import "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"

type suites struct {
	hash hash.Hash
	list []Suite
}

func createSuites(
	hash hash.Hash,
	list []Suite,
) Suites {
	out := suites{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *suites) Hash() hash.Hash {
	return obj.hash
}

// List returns the suites
func (obj *suites) List() []Suite {
	return obj.list
}
