package metadatas

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	name        string
	description string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	return &builder{
		hashAdapter: hashAdapter,
		path:        nil,
		name:        "",
		description: "",
	}
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

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// Now builds a new MetaData instance
func (app *builder) Now() (MetaData, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a MetaData instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a MetaData instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a MetaData instance")
	}

	path := filepath.Join(app.path...)
	data := [][]byte{
		[]byte(path),
		[]byte(app.name),
		[]byte(app.description),
	}

	metaHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createMetadata(*metaHash, app.path, app.name, app.description), nil
}
