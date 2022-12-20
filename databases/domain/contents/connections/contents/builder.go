package contents

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	pHash *hash.Hash
	data  []byte
	pKind *uint
}

func createBuilder() Builder {
	out := builder{
		pHash: nil,
		data:  nil,
		pKind: nil,
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

// WithData adds data to the builder
func (app *builder) WithData(data []byte) Builder {
	app.data = data
	return app
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind uint) Builder {
	app.pKind = &kind
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Content instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Content instance")
	}

	if app.data != nil && len(app.data) <= 0 {
		app.data = nil
	}

	if app.data == nil {
		return nil, errors.New("the data is mandatory in order to build a Content instance")
	}

	return createContent(*app.pHash, app.data, *app.pKind), nil
}
