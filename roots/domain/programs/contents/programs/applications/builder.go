package applications

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	pHash       *hash.Hash
	pIndex      *uint
	pModule     *uint
	attachments Attachments
}

func createBuilder() Builder {
	out := builder{
		pHash:       nil,
		pIndex:      nil,
		pModule:     nil,
		attachments: nil,
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

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithModule adds a module to the builder
func (app *builder) WithModule(module uint) Builder {
	app.pModule = &module
	return app
}

// WithAttachments add attachments to the builder
func (app *builder) WithAttachments(attachments Attachments) Builder {
	app.attachments = attachments
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Application instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Application instance")
	}

	if app.pModule == nil {
		return nil, errors.New("the module is mandatory in order to build an Application instance")
	}

	if app.attachments != nil {
		return createApplicationWithAttachments(*app.pHash, *app.pIndex, *app.pModule, app.attachments), nil
	}

	return createApplication(*app.pHash, *app.pIndex, *app.pModule), nil
}
