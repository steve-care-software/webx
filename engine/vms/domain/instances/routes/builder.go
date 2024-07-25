package routes

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
)

type builder struct {
	hashAdapter hash.Adapter
	tokens      tokens.Tokens
	layer       hash.Hash
	global      omissions.Omission
	token       omissions.Omission
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		tokens:      nil,
		layer:       nil,
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

// WithLayer add layer to the builder
func (app *builder) WithLayer(layer hash.Hash) Builder {
	app.layer = layer
	return app
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens tokens.Tokens) Builder {
	app.tokens = tokens
	return app
}

// WithGlobal add global to the builder
func (app *builder) WithGlobal(global omissions.Omission) Builder {
	app.global = global
	return app
}

// WithToken add token to the builder
func (app *builder) WithToken(token omissions.Omission) Builder {
	app.token = token
	return app
}

// Now builds a new Route instance
func (app *builder) Now() (Route, error) {
	if app.tokens != nil {
		return nil, errors.New("the tokens is mandatory in order to build a Route instance")
	}

	if app.layer != nil {
		return nil, errors.New("the layer is mandatory in order to build a Route instance")
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
		return createRouteWithGlobalAndToken(*pHash, app.layer, app.tokens, app.global, app.token), nil
	}

	if app.global != nil {
		return createRouteWithGlobal(*pHash, app.layer, app.tokens, app.global), nil
	}

	if app.token != nil {
		return createRouteWithToken(*pHash, app.layer, app.tokens, app.token), nil
	}

	return createRoute(*pHash, app.layer, app.tokens), nil
}
