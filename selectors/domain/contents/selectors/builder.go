package selectors

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash   *hash.Hash
	pToken  *hash.Hash
	pInside *hash.Hash
	pFn     *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:   nil,
		pToken:  nil,
		pInside: nil,
		pFn:     nil,
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

// WithToken add token to the builder
func (app *builder) WithToken(token hash.Hash) Builder {
	app.pToken = &token
	return app
}

// WithInside add inside to the builder
func (app *builder) WithInside(inside hash.Hash) Builder {
	app.pInside = &inside
	return app
}

// WithFunc add fn to the builder
func (app *builder) WithFunc(fn hash.Hash) Builder {
	app.pFn = &fn
	return app
}

// Now builds a new Selector instance
func (app *builder) Now() (Selector, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Selector instance")
	}

	if app.pToken == nil {
		return nil, errors.New("the token is mandatory in order to build a Selector instance")
	}

	if app.pInside == nil {
		return nil, errors.New("the inside is mandatory in order to build a Selector instance")
	}

	if app.pFn == nil {
		return nil, errors.New("the func is mandatory in order to build a Selector instance")
	}

	return createSelector(*app.pHash, *app.pToken, *app.pInside, *app.pFn), nil
}
