package suites

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type suiteBuilder struct {
	pHash   *hash.Hash
	isValid bool
	content []byte
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		pHash:   nil,
		isValid: false,
		content: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder()
}

// WithHash adds an hash to the builder
func (app *suiteBuilder) WithHash(hash hash.Hash) SuiteBuilder {
	app.pHash = &hash
	return app
}

// WithContent adds content to the builder
func (app *suiteBuilder) WithContent(content []byte) SuiteBuilder {
	app.content = content
	return app
}

// IsValid flags the builder as valid
func (app *suiteBuilder) IsValid() SuiteBuilder {
	app.isValid = true
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
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
