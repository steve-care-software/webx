package schemas

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/routers/routes/schemas/elements"
	"github.com/steve-care-software/webx/domain/grammars"
)

type schema struct {
	hash     hash.Hash
	grammar  grammars.Grammar
	elements elements.Elements
}

func createSchema(
	hash hash.Hash,
	grammar grammars.Grammar,
	elements elements.Elements,
) Schema {
	out := schema{
		hash:     hash,
		grammar:  grammar,
		elements: elements,
	}

	return &out
}

// Hash returns the hash
func (obj *schema) Hash() hash.Hash {
	return obj.hash
}

// Grammar returns the grammar
func (obj *schema) Grammar() grammars.Grammar {
	return obj.grammar
}

// Elements returns the elements
func (obj *schema) Elements() elements.Elements {
	return obj.elements
}
