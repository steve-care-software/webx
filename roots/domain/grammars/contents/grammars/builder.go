package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type builder struct {
	pHash    *hash.Hash
	pRoot    *hash.Hash
	channels []hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:    nil,
		pRoot:    nil,
		channels: nil,
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

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root hash.Hash) Builder {
	app.pRoot = &root
	return app
}

// WithChannels add channels to the builder
func (app *builder) WithChannels(channels []hash.Hash) Builder {
	app.channels = channels
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Grammar instance")
	}

	if app.pRoot == nil {
		return nil, errors.New("the root is mandatory in order to build a Grammar instance")
	}

	if app.channels != nil && len(app.channels) <= 0 {
		app.channels = nil
	}

	if app.channels != nil {
		return createGrammarWithChannels(*app.pHash, *app.pRoot, app.channels), nil
	}

	return createGrammar(*app.pHash, *app.pRoot), nil
}
