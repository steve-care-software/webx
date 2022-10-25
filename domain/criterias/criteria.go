package criterias

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type criteria struct {
	hash            hash.Hash
	name            string
	includeChannels bool
	child           Criteria
	pIndex          *uint
}

func createCriteria(
	hash hash.Hash,
	name string,
	includeChannels bool,
) Criteria {
	return createCriteriaInternally(hash, name, includeChannels, nil, nil)
}

func createCriteriaWithChild(
	hash hash.Hash,
	name string,
	includeChannels bool,
	child Criteria,
) Criteria {
	return createCriteriaInternally(hash, name, includeChannels, child, nil)
}

func createCriteriaWithIndex(
	hash hash.Hash,
	name string,
	includeChannels bool,
	pIndex *uint,
) Criteria {
	return createCriteriaInternally(hash, name, includeChannels, nil, pIndex)
}

func createCriteriaWithChildAndIndex(
	hash hash.Hash,
	name string,
	includeChannels bool,
	child Criteria,
	pIndex *uint,
) Criteria {
	return createCriteriaInternally(hash, name, includeChannels, child, pIndex)
}

func createCriteriaInternally(
	hash hash.Hash,
	name string,
	includeChannels bool,
	child Criteria,
	pIndex *uint,
) Criteria {
	out := criteria{
		hash:            hash,
		name:            name,
		includeChannels: includeChannels,
		child:           child,
		pIndex:          pIndex,
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

// HasIndex returns true if there is an index, false otherwise
func (obj *criteria) HasIndex() bool {
	return obj.pIndex != nil
}

// Index returns the index
func (obj *criteria) Index() *uint {
	return obj.pIndex
}
