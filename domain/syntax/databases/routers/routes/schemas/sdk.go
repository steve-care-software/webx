package schemas

import (
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/routers/routes/schemas/elements"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a schema builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithElements(elements elements.Elements) Builder
	Now() (Schema, error)
}

// Schema represents a route schema
type Schema interface {
	Hash() hash.Hash
	Grammar() grammars.Grammar
	Elements() elements.Elements
}
