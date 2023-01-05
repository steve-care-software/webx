package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type lineBuilder struct {
	hashAdapter hash.Adapter
	containers  []Container
}

func createLineBuilder(
	hashAdapter hash.Adapter,
) LineBuilder {
	out := lineBuilder{
		hashAdapter: hashAdapter,
		containers:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder(
		app.hashAdapter,
	)
}

// WithContainers add containers to the builder
func (app *lineBuilder) WithContainers(containers []Container) LineBuilder {
	app.containers = containers
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.containers != nil && len(app.containers) <= 0 {
		app.containers = nil
	}

	if app.containers == nil {
		return nil, errors.New("there must be at least 1 Container in order to build a Line instance")
	}

	data := [][]byte{}
	for _, oneContainer := range app.containers {
		data = append(data, oneContainer.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createLine(*pHash, app.containers), nil
}
