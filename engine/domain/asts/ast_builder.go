package asts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type astAstBuilder struct {
	hashAdapter hash.Adapter
	library     NFTs
	entry       hash.Hash
}

func createAstBuilder(
	hashAdapter hash.Adapter,
) AstBuilder {
	out := astAstBuilder{
		hashAdapter: hashAdapter,
		library:     nil,
		entry:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *astAstBuilder) Create() AstBuilder {
	return createAstBuilder(
		app.hashAdapter,
	)
}

// WithLibrary adds a library to the builder
func (app *astAstBuilder) WithLibrary(library NFTs) AstBuilder {
	app.library = library
	return app
}

// WithEntry adds an entry to the builder
func (app *astAstBuilder) WithEntry(entry hash.Hash) AstBuilder {
	app.entry = entry
	return app
}

// Now builds a new AST instance
func (app *astAstBuilder) Now() (AST, error) {
	if app.library == nil {
		return nil, errors.New("the library is mandatory in order to build an AST instance")
	}

	if app.entry == nil {
		return nil, errors.New("the entry hash is mandatory in order to build an AST instance")
	}

	return createAST(
		app.library,
		app.entry,
	), nil
}
