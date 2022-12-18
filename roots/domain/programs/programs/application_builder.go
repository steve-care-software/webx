package programs

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/programs/programs/modules"
)

type applicationBuilder struct {
	hashAdapter hash.Adapter
	pIndex      *uint
	module      modules.Module
	attachments Attachments
}

func createApplicationBuilder(
	hashAdapter hash.Adapter,
) ApplicationBuilder {
	out := applicationBuilder{
		hashAdapter: hashAdapter,
		pIndex:      nil,
		module:      nil,
		attachments: nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() ApplicationBuilder {
	return createApplicationBuilder(
		app.hashAdapter,
	)
}

// WithIndex adds an index to the builder
func (app *applicationBuilder) WithIndex(index uint) ApplicationBuilder {
	app.pIndex = &index
	return app
}

// WithModule adds a module to the builder
func (app *applicationBuilder) WithModule(module modules.Module) ApplicationBuilder {
	app.module = module
	return app
}

// WithAttachments add attachments to the builder
func (app *applicationBuilder) WithAttachments(attachments Attachments) ApplicationBuilder {
	app.attachments = attachments
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Application instance")
	}

	if app.module == nil {
		return nil, errors.New("the module is mandatory in order to build an Application instance")
	}

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", *app.pIndex)),
		[]byte(fmt.Sprintf("%d", app.module.Index())),
	}

	if app.attachments != nil {
		data = append(data, app.attachments.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.attachments != nil {
		return createApplicationWithAttachments(*pHash, *app.pIndex, app.module, app.attachments), nil
	}

	return createApplication(*pHash, *app.pIndex, app.module), nil
}
