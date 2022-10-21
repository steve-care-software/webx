package indexes

import (
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type index struct {
	hash     hash.Hash
	name     string
	criteria criterias.Criteria
}

func createIndex(
	hash hash.Hash,
	name string,
	criteria criterias.Criteria,
) Index {
	out := index{
		hash:     hash,
		name:     name,
		criteria: criteria,
	}

	return &out
}

// Hash returns the hash
func (obj *index) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *index) Name() string {
	return obj.name
}

// Criteria returns the criteria
func (obj *index) Criteria() criterias.Criteria {
	return obj.criteria
}
