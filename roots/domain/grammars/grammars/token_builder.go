package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type tokenBuilder struct {
	hashAdapter hash.Adapter
	name        string
	block       Block
	suites      Suites
}

func createTokenBuilder(
	hashAdapter hash.Adapter,
) TokenBuilder {
	out := tokenBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		block:       nil,
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// WithBlock adds a block to the builder
func (app *tokenBuilder) WithBlock(block Block) TokenBuilder {
	app.block = block
	return app
}

// WithSuites add suites to the builder
func (app *tokenBuilder) WithSuites(suites Suites) TokenBuilder {
	app.suites = suites
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.block == nil {
		return nil, errors.New("the block is mandatory in order to build a Token instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.block.Hash().Bytes(),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createTokenWithSuites(*pHash, app.name, app.block, app.suites), nil
	}

	return createToken(*pHash, app.name, app.block), nil
}
