package grammars

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/grammars/values"
)

type composeElement struct {
	hash       hash.Hash
	value      values.Value
	occurences uint
}

func createComposeElement(
	hash hash.Hash,
	value values.Value,
	occurences uint,
) ComposeElement {
	out := composeElement{
		hash:       hash,
		value:      value,
		occurences: occurences,
	}

	return &out
}

// Hash returns the hash
func (obj *composeElement) Hash() hash.Hash {
	return obj.hash
}

// Points returns the points
func (obj *composeElement) Points() uint {
	return obj.occurences
}

// Value returns the value
func (obj *composeElement) Value() values.Value {
	return obj.value
}

// Occurences returns the occurences
func (obj *composeElement) Occurences() uint {
	return obj.occurences
}
