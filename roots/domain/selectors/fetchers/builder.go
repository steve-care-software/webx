package fetchers

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	pHash      *hash.Hash
	pRecursive *hash.Hash
	pSelector  *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:      nil,
		pRecursive: nil,
		pSelector:  nil,
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

// WithRecursive adds a recursive to the builder
func (app *builder) WithRecursive(recursive hash.Hash) Builder {
	app.pRecursive = &recursive
	return app
}

// WithSelector adds a selector to the builder
func (app *builder) WithSelector(selector hash.Hash) Builder {
	app.pSelector = &selector
	return app
}

// Now builds a new Fetcher instance
func (app *builder) Now() (Fetcher, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Fetcher instance")
	}

	if app.pRecursive != nil {
		content := createContentWithRecursive(app.pRecursive)
		return createFetcher(*app.pHash, content), nil
	}

	if app.pSelector != nil {
		content := createContentWithSelector(app.pSelector)
		return createFetcher(*app.pHash, content), nil
	}

	return nil, errors.New("the Fetcher is invalid")
}
