package databases

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        string
	description string
	head        string
	isActive    string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        "",
		description: "",
		head:        "",
		isActive:    "",
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
func (app *builder) WithPath(path string) Builder {
	app.path = path
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithHead adds a head to the builder
func (app *builder) WithHead(head string) Builder {
	app.head = head
	return app
}

// WithActive returns the isActive
func (app *builder) WithActive(isActive string) Builder {
	app.isActive = isActive
	return app
}

// Now builds a new Database
func (app *builder) Now() (Database, error) {
	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build a Database instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Database instance")
	}

	if app.head == "" {
		return nil, errors.New("the head is mandatory in order to build a Database instance")
	}

	if app.isActive == "" {
		return nil, errors.New("the isActive is mandatory in order to build a Database instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.path),
		[]byte(app.description),
		[]byte(app.head),
		[]byte(app.isActive),
	})

	if err != nil {
		return nil, err
	}

	return createDatabase(*pHash, app.path, app.description, app.head, app.isActive), nil
}
