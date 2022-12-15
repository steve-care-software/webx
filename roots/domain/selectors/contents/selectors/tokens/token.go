package tokens

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type token struct {
	hash     hash.Hash
	reverse  hash.Hash
	element  Element
	pContent *uint
}

func createToken(
	hash hash.Hash,
	reverse hash.Hash,
	element Element,
) Token {
	return createTokenInternally(hash, reverse, element, nil)
}

func createTokenWithContentIndex(
	hash hash.Hash,
	reverse hash.Hash,
	element Element,
	pContent *uint,
) Token {
	return createTokenInternally(hash, reverse, element, pContent)
}

func createTokenInternally(
	hash hash.Hash,
	reverse hash.Hash,
	element Element,
	pContent *uint,
) Token {
	out := token{
		hash:     hash,
		reverse:  reverse,
		element:  element,
		pContent: pContent,
	}

	return &out
}

// Hash returns the hash
func (obj *token) Hash() hash.Hash {
	return obj.hash
}

// Reverse returns the reverse
func (obj *token) Reverse() hash.Hash {
	return obj.reverse
}

// Element returns the element
func (obj *token) Element() Element {
	return obj.element
}

// HasContent returns true if there is content, false otherwise
func (obj *token) HasContent() bool {
	return obj.pContent != nil
}

// Content returns the content, if any
func (obj *token) Content() *uint {
	return obj.pContent
}
