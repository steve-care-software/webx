package conditions

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

type comparisons struct {
	hash hash.Hash
	list []Comparison
}

func createComparisons(
	hash hash.Hash,
	list []Comparison,
) Comparisons {
	out := comparisons{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *comparisons) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *comparisons) List() []Comparison {
	return obj.list
}
