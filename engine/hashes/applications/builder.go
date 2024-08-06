package applications

import (
	"errors"

	bytes_applications "github.com/steve-care-software/webx/engine/bytes/applications"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/hashes/domain/pointers"
)

type builder struct {
	pointerAdapter   pointers.Adapter
	pointersBuilder  pointers.Builder
	pointerBuilder   pointers.PointerBuilder
	delimiterBuilder delimiters.DelimiterBuilder
	bytesApp         bytes_applications.Application
}

func createBuilder(
	pointerAdapter pointers.Adapter,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	delimiterBuilder delimiters.DelimiterBuilder,
) Builder {
	out := builder{
		pointerAdapter:   pointerAdapter,
		pointersBuilder:  pointersBuilder,
		pointerBuilder:   pointerBuilder,
		delimiterBuilder: delimiterBuilder,
		bytesApp:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.pointerAdapter,
		app.pointersBuilder,
		app.pointerBuilder,
		app.delimiterBuilder,
	)
}

// WithBytes adds a bytes application to the builder
func (app *builder) WithBytes(bytesApp bytes_applications.Application) Builder {
	app.bytesApp = bytesApp
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.bytesApp == nil {
		return nil, errors.New("the bytes application is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.bytesApp,
		app.pointerAdapter,
		app.pointersBuilder,
		app.pointerBuilder,
		app.delimiterBuilder,
	), nil
}
