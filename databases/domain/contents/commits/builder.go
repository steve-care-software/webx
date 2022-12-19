package commits

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type builder struct {
	pHash   *hash.Hash
	values  hashtrees.HashTree
	pParent *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:   nil,
		values:  nil,
		pParent: nil,
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

// WithValues add values to the builder
func (app *builder) WithValues(values hashtrees.HashTree) Builder {
	app.values = values
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.pParent = &parent
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Commit instance")
	}

	if app.values == nil {
		return nil, errors.New("the values is mandatory in order to build a Commit instance")
	}

	return createCommit(*app.pHash, app.values), nil
}
