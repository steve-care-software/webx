package suites

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash   *hash.Hash
	isValid bool
	content []byte
}

func createBuilder() Builder {
	out := builder{
		pHash:   nil,
		isValid: false,
		content: nil,
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

// WithContent adds content to the builder
func (app *builder) WithContent(content []byte) Builder {
	app.content = content
	return app
}

// IsValid flags the builder as valid
func (app *builder) IsValid() Builder {
	app.isValid = true
	return app
}

// Now builds a new Suite instance
func (app *builder) Now() (Suite, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Suite instance")
	}

	if app.content != nil && len(app.content) <= 0 {
		app.content = nil
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Suite instance")
	}

	return createSuite(*app.pHash, app.isValid, app.content), nil
}
