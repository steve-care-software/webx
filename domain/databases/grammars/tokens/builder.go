package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type builder struct {
	pHash  *hash.Hash
	lines  Lines
	suites Suites
}

func createBuilder() Builder {
	out := builder{
		pHash:  nil,
		lines:  nil,
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

// WithLines add lines to the builder
func (app *builder) WithLines(lines Lines) Builder {
	app.lines = lines
	return app
}

// WithSuites add suites to the builder
func (app *builder) WithSuites(suites Suites) Builder {
	app.suites = suites
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Token instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Token instance")
	}

	if app.suites != nil {
		return createTokenWithSuites(*app.pHash, app.lines, app.suites), nil
	}

	return createToken(*app.pHash, app.lines), nil
}
