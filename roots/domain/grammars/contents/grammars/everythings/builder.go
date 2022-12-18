package everythings

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	pHash      *hash.Hash
	pException *hash.Hash
	pEscape    *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:      nil,
		pException: nil,
		pEscape:    nil,
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

// WithException adds an exception to the builder
func (app *builder) WithException(exception hash.Hash) Builder {
	app.pException = &exception
	return app
}

// WithEscape adds an escape to the builder
func (app *builder) WithEscape(escape hash.Hash) Builder {
	app.pEscape = &escape
	return app
}

// Now builds a new Everything instance
func (app *builder) Now() (Everything, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Everything instance")
	}

	if app.pException == nil {
		return nil, errors.New("the exception is mandatory in order to build an Everything instance")
	}

	if app.pEscape != nil {
		return createEverythingWithEscape(*app.pHash, *app.pException, app.pEscape), nil
	}

	return createEverything(*app.pHash, *app.pException), nil
}
