package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hashtrees"
)

type builder struct {
	pHash   *hash.Hash
	name    string
	history hashtrees.HashTree
}

func createBuilder() Builder {
	out := builder{
		pHash:   nil,
		name:    "",
		history: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.pHash = &hash
	return app
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
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Grammar instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Grammar instance")
	}

	if app.history != nil {
		return createGrammarWithHistory(*app.pHash, app.name, app.history), nil
	}

	return createGrammar(*app.pHash, app.name), nil
}
