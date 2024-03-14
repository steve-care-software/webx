package libraries

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases"
)

type builder struct {
	hashAdapter hash.Adapter
	compiler    compilers.Compiler
	database    databases.Database
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		compiler:    nil,
		database:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCompiler adds a compiler to the builder
func (app *builder) WithCompiler(compiler compilers.Compiler) Builder {
	app.compiler = compiler
	return app
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(database databases.Database) Builder {
	app.database = database
	return app
}

// Now builds a new Library instance
func (app *builder) Now() (Library, error) {
	data := [][]byte{}
	if app.compiler != nil {
		data = append(data, []byte("compiler"))
		data = append(data, app.compiler.Hash().Bytes())
	}

	if app.database != nil {
		data = append(data, []byte("database"))
		data = append(data, app.database.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Library is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.compiler != nil {
		return createLibraryWithCompiler(*pHash, app.compiler), nil
	}

	return createLibraryWithDatabase(*pHash, app.database), nil
}
