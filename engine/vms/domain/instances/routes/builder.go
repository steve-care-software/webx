package routes

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	tokens      Tokens
	global      Omission
	token       Omission
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		tokens:      nil,
		global:      nil,
		token:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens Tokens) Builder {
	app.tokens = tokens
	return app
}

// WithGlobal add global to the builder
func (app *builder) WithGlobal(global Omission) Builder {
	app.global = global
	return app
}

// WithToken add token to the builder
func (app *builder) WithToken(token Omission) Builder {
	app.token = token
	return app
}

// Now builds a new Route instance
func (app *builder) Now() (Route, error) {
	if app.tokens != nil {
		return nil, errors.New("the tokens is mandatory in order to build a Route instance")
	}

	data := [][]byte{
		[]byte("token"),
		app.tokens.Hash().Bytes(),
	}

	if app.global != nil {
		data = append(data, app.global.Hash().Bytes())
	}

	if app.tokens != nil {
		data = append(data, app.tokens.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.global != nil && app.token != nil {
		return createRouteWithGlobalAndToken(*pHash, app.tokens, app.global, app.token), nil
	}

	if app.global != nil {
		return createRouteWithGlobal(*pHash, app.tokens, app.global), nil
	}

	if app.token != nil {
		return createRouteWithToken(*pHash, app.tokens, app.token), nil
	}

	return createRoute(*pHash, app.tokens), nil
}
