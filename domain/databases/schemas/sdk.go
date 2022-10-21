package schemas

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/routers"
	"github.com/steve-care-software/webx/domain/databases/schemas/indexes"
	"github.com/steve-care-software/webx/domain/grammars"
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
	WithRouter(router routers.Router) Builder
	WithIndexes(indexes indexes.Indexes) Builder
	Now() (Schema, error)
}

// Schema represents a schema
type Schema interface {
	Hash() hash.Hash
	Grammar() grammars.Grammar
	Router() routers.Router
	HasIndexes() bool
	Indexes() indexes.Indexes
}
