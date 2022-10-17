package routes

import (
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/programs"
	"github.com/steve-care-software/syntax/domain/syntax/databases/routers/routes/schemas"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a route builder
type Builder interface {
	Create() Builder
	WithSchema(schema schemas.Schema) Builder
	WithProgram(program programs.Program) Builder
	Now() (Route, error)
}

// Route represents a route
type Route interface {
	Hash() hash.Hash
	Schema() schemas.Schema
	Program() programs.Program
}
