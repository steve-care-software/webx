package routes

import (
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/programs"
	"github.com/steve-care-software/syntax/domain/syntax/databases/routers/routes/schemas"
)

type route struct {
	hash    hash.Hash
	schema  schemas.Schema
	program programs.Program
}

func createRoute(
	hash hash.Hash,
	schema schemas.Schema,
	program programs.Program,
) Route {
	out := route{
		hash:    hash,
		schema:  schema,
		program: program,
	}

	return &out
}

// Hash returns the hash
func (obj *route) Hash() hash.Hash {
	return obj.hash
}

// Schema returns the schema
func (obj *route) Schema() schemas.Schema {
	return obj.schema
}

// Program returns the program
func (obj *route) Program() programs.Program {
	return obj.program
}
