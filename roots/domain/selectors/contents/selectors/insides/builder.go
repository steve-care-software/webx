package insides

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	pHash    *hash.Hash
	pFn      *hash.Hash
	fetchers []hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:    nil,
		pFn:      nil,
		fetchers: nil,
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

// WithFn adds a fn to the builder
func (app *builder) WithFn(fn hash.Hash) Builder {
	app.pFn = &fn
	return app
}

// WithFetchers adds a fetchers to the builder
func (app *builder) WithFetchers(fetchers []hash.Hash) Builder {
	app.fetchers = fetchers
	return app
}

// Now builds a new Inside instance
func (app *builder) Now() (Inside, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mndatory in order to build an Inside instance")
	}

	if app.fetchers != nil && len(app.fetchers) <= 0 {
		app.fetchers = nil
	}

	if app.pFn != nil {
		content := createContentWithFn(*app.pFn)
		return createInside(*app.pHash, content), nil
	}

	if app.fetchers != nil {
		content := createContentWithFetchers(app.fetchers)
		return createInside(*app.pHash, content), nil
	}

	return nil, errors.New("the Inside is invalid")
}
