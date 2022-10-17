package criterias

import "github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"

type criteria struct {
	hash            hash.Hash
	name            string
	index           uint
	includeChannels bool
	child           Criteria
}

func createCriteria(
	hash hash.Hash,
	name string,
	index uint,
	includeChannels bool,
) Criteria {
	return createCriteriaInternally(hash, name, index, includeChannels, nil)
}

func createCriteriaWithChild(
	hash hash.Hash,
	name string,
	index uint,
	includeChannels bool,
	child Criteria,
) Criteria {
	return createCriteriaInternally(hash, name, index, includeChannels, child)
}

func createCriteriaInternally(
	hash hash.Hash,
	name string,
	index uint,
	includeChannels bool,
	child Criteria,
) Criteria {
	out := criteria{
		hash:            hash,
		name:            name,
		index:           index,
		includeChannels: includeChannels,
		child:           child,
	}

	return &out
}

// Hash returns the hash
func (obj *criteria) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *criteria) Name() string {
	return obj.name
}

// Index returns the index
func (obj *criteria) Index() uint {
	return obj.index
}

// IncludeChannels returns true if channels are included, false otherwise
func (obj *criteria) IncludeChannels() bool {
	return obj.includeChannels
}

// HasChild returns true if there is a child, false otherwise
func (obj *criteria) HasChild() bool {
	return obj.child != nil
}

// Child returns the child, if any
func (obj *criteria) Child() Criteria {
	return obj.child
}
