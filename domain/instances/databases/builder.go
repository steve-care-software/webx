package databases

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	description string
	head        commits.Commit
	isActive    bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        nil,
		description: "",
		head:        nil,
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

// WithHead adds a head to the builder
func (app *builder) WithHead(head commits.Commit) Builder {
	app.head = head
	return app
}

// IsActive flags the builder as active
func (app *builder) IsActive() Builder {
	app.isActive = true
	return app
}

// Now builds a new Database
func (app *builder) Now() (Database, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Database instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Database instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Database instance")
	}

	isActive := "false"
	if app.isActive {
		isActive = "true"
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		[]byte(app.description),
		app.head.Hash().Bytes(),
		[]byte(isActive),
	})

	if err != nil {
		return nil, err
	}

	return createDatabase(*pHash, app.path, app.description, app.head, app.isActive), nil
}
