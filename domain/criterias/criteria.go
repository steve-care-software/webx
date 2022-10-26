package criterias

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type criteria struct {
	hash    hash.Hash
	current Tail
	next    Node
}

func createCriteria(
	hash hash.Hash,
	current Tail,
) Criteria {
	return createCriteriaInternally(hash, current, nil)
}

func createCriteriaWithNext(
	hash hash.Hash,
	current Tail,
	next Node,
) Criteria {
	return createCriteriaInternally(hash, current, next)
}

func createCriteriaInternally(
	hash hash.Hash,
	current Tail,
	next Node,
) Criteria {
	out := criteria{
		hash:    hash,
		current: current,
		next:    next,
	}

	return &out
}

// Hash retruns the hash
func (obj *criteria) Hash() hash.Hash {
	return obj.hash
}

// Current returns the current tail
func (obj *criteria) Current() Tail {
	return obj.current
}

// HasNext returns true if there is a next node, false otherwise
func (obj *criteria) HasNext() bool {
	return obj.next != nil
}

// Next returns the next node, if any
func (obj *criteria) Next() Node {
	return obj.next
}
