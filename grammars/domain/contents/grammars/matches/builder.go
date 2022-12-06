package matches

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash  *hash.Hash
	pToken *hash.Hash
	suites []hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:  nil,
		pToken: nil,
		suites: nil,
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

// WithSuites add suites to the builder
func (app *builder) WithSuites(suites []hash.Hash) Builder {
	app.suites = suites
	return app
}

// Now builds a new Match instance
func (app *builder) Now() (Match, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Match instance")
	}

	if app.pToken == nil {
		return nil, errors.New("the token hash is mandatory in order to build a Match instance")
	}

	if app.suites != nil && len(app.suites) <= 0 {
		app.suites = nil
	}

	if app.suites == nil {
		return nil, errors.New("there must be at least 1 suite hash in order to build a Match instance")
	}

	return createMatch(*app.pHash, *app.pToken, app.suites), nil
}
