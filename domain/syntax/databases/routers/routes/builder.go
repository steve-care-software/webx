package routes

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/programs"
	"github.com/steve-care-software/syntax/domain/syntax/databases/routers/routes/schemas"
)

type builder struct {
	hashAdapter hash.Adapter
	schema      schemas.Schema
	program     programs.Program
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		schema:      nil,
		program:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithSchema adds a schema to the builder
func (app *builder) WithSchema(schema schemas.Schema) Builder {
	app.schema = schema
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program programs.Program) Builder {
	app.program = program
	return app
}

// Now builds a new Route instance
func (app *builder) Now() (Route, error) {
	if app.schema == nil {
		return nil, errors.New("the schema is mandatory in order to build a Route instance")
	}

	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build a Route instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.schema.Hash().Bytes(),
		app.program.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createRoute(*pHash, app.schema, app.program), nil
}
