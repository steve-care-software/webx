package executions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	executable  string
	content     Content
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		executable:  "",
		content:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithExecutable adds an executable to the builder
func (app *builder) WithExecutable(executable string) Builder {
	app.executable = executable
	return app
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	if app.executable == "" {
		return nil, errors.New("the executable is mandatory in order to build an Execution instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Execution instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.executable),
		app.content.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createExecution(*pHash, app.executable, app.content), nil
}
