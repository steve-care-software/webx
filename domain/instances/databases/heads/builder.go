package heads

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	description string
	isActive    bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        nil,
		description: "",
		isActive:    false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// IsActive flags the builder as active
func (app *builder) IsActive() Builder {
	app.isActive = true
	return app
}

// Now builds a new Head instance
func (app *builder) Now() (Head, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build an Head instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build an Head instance")
	}

	isActive := "false"
	if app.isActive {
		isActive = "true"
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		[]byte(isActive),
		[]byte(app.description),
	})

	if err != nil {
		return nil, err
	}

	return createHead(*pHash, app.path, app.description, app.isActive), nil
}
