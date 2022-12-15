package elements

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type element struct {
	hash        hash.Hash
	entity      entities.Entity
	cardinality Cardinality
	content     Content
}

func createElement(
	hash hash.Hash,
	cardinality Cardinality,
	content Content,
) Element {
	out := element{
		hash:        hash,
		cardinality: cardinality,
		content:     content,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Cardinality returns the cardinality
func (obj *element) Cardinality() Cardinality {
	return obj.cardinality
}

// Content returns the content
func (obj *element) Content() Content {
	return obj.content
}
