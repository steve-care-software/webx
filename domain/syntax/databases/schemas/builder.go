package schemas

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/routers"
	"github.com/steve-care-software/syntax/domain/syntax/databases/schemas/indexes"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

type builder struct {
	hashAdapter hash.Adapter
	grammar     grammars.Grammar
	router      routers.Router
	indexes     indexes.Indexes
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		grammar:     nil,
		router:      nil,
		indexes:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar grammars.Grammar) Builder {
	app.grammar = grammar
	return app
}

// WithRouter adds a router to the builder
func (app *builder) WithRouter(router routers.Router) Builder {
	app.router = router
	return app
}

// WithIndexes add indexes to the builder
func (app *builder) WithIndexes(indexes indexes.Indexes) Builder {
	app.indexes = indexes
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Schema instance")
	}

	if app.router == nil {
		return nil, errors.New("the router is mandatory in order to build a Schema instance")
	}

	data := [][]byte{
		app.grammar.Hash().Bytes(),
		app.router.Hash().Bytes(),
	}

	if app.indexes != nil {
		data = append(data, app.indexes.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.indexes != nil {
		return createSchemaWithIndexes(*pHash, app.grammar, app.router, app.indexes), nil
	}

	return createSchema(*pHash, app.grammar, app.router), nil
}
