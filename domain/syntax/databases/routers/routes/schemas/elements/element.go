package elements

import (
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

type element struct {
	hash     hash.Hash
	name     string
	criteria criterias.Criteria
}

func createElement(
	hash hash.Hash,
	name string,
	criteria criterias.Criteria,
) Element {
	out := element{
		hash:     hash,
		name:     name,
		criteria: criteria,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *element) Name() string {
	return obj.name
}

// Criteria returns the criteria
func (obj *element) Criteria() criterias.Criteria {
	return obj.criteria
}
