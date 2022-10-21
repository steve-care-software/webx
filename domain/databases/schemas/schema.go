package schemas

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/routers"
	"github.com/steve-care-software/webx/domain/databases/schemas/indexes"
	"github.com/steve-care-software/webx/domain/grammars"
)

type schema struct {
	hash    hash.Hash
	grammar grammars.Grammar
	router  routers.Router
	indexes indexes.Indexes
}

func createSchema(
	hash hash.Hash,
	grammar grammars.Grammar,
	router routers.Router,
) Schema {
	return createSchemaInternally(hash, grammar, router, nil)
}

func createSchemaWithIndexes(
	hash hash.Hash,
	grammar grammars.Grammar,
	router routers.Router,
	indexes indexes.Indexes,
) Schema {
	return createSchemaInternally(hash, grammar, router, indexes)
}

func createSchemaInternally(
	hash hash.Hash,
	grammar grammars.Grammar,
	router routers.Router,
	indexes indexes.Indexes,
) Schema {
	out := schema{
		hash:    hash,
		grammar: grammar,
		router:  router,
		indexes: indexes,
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

// Router returns the router
func (obj *schema) Router() routers.Router {
	return obj.router
}

// HasIndexes returns true if there is indexes, false otherwise
func (obj *schema) HasIndexes() bool {
	return obj.indexes != nil
}

// Indexes returns the indexes, if any
func (obj *schema) Indexes() indexes.Indexes {
	return obj.indexes
}
