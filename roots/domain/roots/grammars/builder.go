package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	history     hashtrees.HashTree
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		history:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithHistory adds an history to the builder
func (app *builder) WithHistory(history hashtrees.HashTree) Builder {
	app.history = history
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Grammar instance")
	}

	data := [][]byte{
		[]byte(app.name),
	}

	if app.history != nil {
		data = append(data, app.history.Head().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.history != nil {
		return createGrammarWithHistory(*pHash, app.name, app.history), nil
	}

	return createGrammar(*pHash, app.name), nil
}
