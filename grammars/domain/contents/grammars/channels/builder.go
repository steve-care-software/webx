package channels

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash  *hash.Hash
	pToken *hash.Hash
	pPrev  *hash.Hash
	pNext  *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:  nil,
		pToken: nil,
		pPrev:  nil,
		pNext:  nil,
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

// WithToken adds a token to the builder
func (app *builder) WithToken(token hash.Hash) Builder {
	app.pToken = &token
	return app
}

// WithPrevious adds a previous to the builder
func (app *builder) WithPrevious(previous hash.Hash) Builder {
	app.pPrev = &previous
	return app
}

// WithNext adds a next to the builder
func (app *builder) WithNext(next hash.Hash) Builder {
	app.pNext = &next
	return app
}

// Now builds a new Channel instance
func (app *builder) Now() (Channel, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Channel instance")
	}

	if app.pToken == nil {
		return nil, errors.New("the token is mandatory in order to build a Channel instance")
	}

	if app.pPrev != nil && app.pNext != nil {
		return createChannelWithPreviousAndNext(*app.pHash, *app.pToken, app.pPrev, app.pNext), nil
	}

	if app.pPrev != nil {
		return createChannelWithPrevious(*app.pHash, *app.pToken, app.pPrev), nil
	}

	if app.pNext != nil {
		return createChannelWithNext(*app.pHash, *app.pToken, app.pNext), nil
	}

	return createChannel(*app.pHash, *app.pToken), nil
}
