package selectors

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type selector struct {
	hash    hash.Hash
	grammar hash.Hash
	token   hash.Hash
	inside  hash.Hash
	fn      hash.Hash
}

func createSelector(
	hash hash.Hash,
	grammar hash.Hash,
	token hash.Hash,
	inside hash.Hash,
	fn hash.Hash,
) Selector {
	out := selector{
		hash:    hash,
		grammar: grammar,
		token:   token,
		inside:  inside,
		fn:      fn,
	}

	return &out
}

// Hash returns the hash
func (obj *selector) Hash() hash.Hash {
	return obj.hash
}

// Grammar returns the grammar hash
func (obj *selector) Grammar() hash.Hash {
	return obj.grammar
}

// Token returns the token
func (obj *selector) Token() hash.Hash {
	return obj.token
}

// Inside returns the inside
func (obj *selector) Inside() hash.Hash {
	return obj.inside
}

// Func returns the func
func (obj *selector) Func() hash.Hash {
	return obj.fn
}
