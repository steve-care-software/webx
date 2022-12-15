package tokens

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

type element struct {
	el    hash.Hash
	index uint
}

func createElement(
	el hash.Hash,
	index uint,
) Element {
	out := element{
		el:    el,
		index: index,
	}

	return &out
}

// Element returns the element
func (obj *element) Element() hash.Hash {
	return obj.el
}

// Index returns the index
func (obj *element) Index() uint {
	return obj.index
}
