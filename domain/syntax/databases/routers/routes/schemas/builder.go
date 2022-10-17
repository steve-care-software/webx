package schemas

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/routers/routes/schemas/elements"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

type builder struct {
	hashAdapter hash.Adapter
	grammar     grammars.Grammar
	elements    elements.Elements
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		grammar:     nil,
		elements:    nil,
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

// WithElements add elements to the builder
func (app *builder) WithElements(elements elements.Elements) Builder {
	app.elements = elements
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.grammar == nil {
		return nil, errors.New("the schema is mandatory in order to build a Schema instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Schema instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.grammar.Hash().Bytes(),
		app.elements.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSchema(*pHash, app.grammar, app.elements), nil
}
