package criterias

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type node struct {
	next Criteria
	tail Tail
}

func createNodeWithNext(
	next Criteria,
) Node {
	return createNodeInternally(next, nil)
}

func createNodeWithTail(
	tail Tail,
) Node {
	return createNodeInternally(nil, tail)
}

func createNodeInternally(
	next Criteria,
	tail Tail,
) Node {
	out := node{
		next: next,
		tail: tail,
	}

	return &out
}

// Hash returns the hash
func (obj *node) Hash() hash.Hash {
	if obj.IsNext() {
		return obj.Next().Hash()
	}

	return obj.Tail().Hash()
}

// IsNext returns true if there is a next criteria, false otherwise
func (obj *node) IsNext() bool {
	return obj.next != nil
}

// Next returns the next criteria, if any
func (obj *node) Next() Criteria {
	return obj.next
}

// IsTail returns true if there is a tail, false otherwise
func (obj *node) IsTail() bool {
	return obj.tail != nil
}

// Tail returns the tail, if any
func (obj *node) Tail() Tail {
	return obj.tail
}
