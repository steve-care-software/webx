package grammars

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/grammars/grammars/cardinalities"
)

type element struct {
	hash        hash.Hash
	content     ElementContent
	cardinality cardinalities.Cardinality
}

func createElement(
	hash hash.Hash,
	content ElementContent,
	cardinality cardinalities.Cardinality,
) Element {
	out := element{
		hash:        hash,
		content:     content,
		cardinality: cardinality,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Points returns the amount of points an element contains
func (obj *element) Points() uint {
	content := obj.Content()
	if content.IsValue() {
		return pointsPerValue
	}

	if content.IsExternal() {
		return content.External().Grammar().Root().Block().Points()
	}

	if content.IsRecursive() {
		return uint(0)
	}

	return content.Instance().Points()
}

// Name returns the name
func (obj *element) Name() string {
	if obj.content.IsValue() {
		return obj.content.Value().Name()
	}

	if obj.content.IsExternal() {
		return obj.content.External().Name()
	}

	if obj.content.IsRecursive() {
		return obj.content.Recursive()
	}

	return obj.content.Instance().Name()
}

// Content returns the content
func (obj *element) Content() ElementContent {
	return obj.content
}

// Cardinality returns the cardinality
func (obj *element) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}
