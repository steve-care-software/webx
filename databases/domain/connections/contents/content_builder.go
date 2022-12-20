package contents

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type contentBuilder struct {
	pHash *hash.Hash
	data  []byte
	pKind *uint
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
		pHash: nil,
		data:  nil,
		pKind: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// WithHash adds an hash to the builder
func (app *contentBuilder) WithHash(hash hash.Hash) ContentBuilder {
	app.pHash = &hash
	return app
}

// WithData adds data to the builder
func (app *contentBuilder) WithData(data []byte) ContentBuilder {
	app.data = data
	return app
}

// WithKind adds a kind to the builder
func (app *contentBuilder) WithKind(kind uint) ContentBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
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
