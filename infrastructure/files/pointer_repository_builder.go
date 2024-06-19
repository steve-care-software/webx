package files

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/pointers"
)

type pointerRepositoryBuilder struct {
	pointersAdapter pointers.Adapter
	pointersBuilder pointers.Builder
	basePath        []string
}

func createPointerRepositoryBuilder(
	pointersAdapter pointers.Adapter,
	pointersBuilder pointers.Builder,
) pointers.RepositoryBuilder {
	out := pointerRepositoryBuilder{
		pointersAdapter: pointersAdapter,
		pointersBuilder: pointersBuilder,
		basePath:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointerRepositoryBuilder) Create() pointers.RepositoryBuilder {
	return createPointerRepositoryBuilder(
		app.pointersAdapter,
		app.pointersBuilder,
	)
}

// WithBasePath adds a basePath to the builder
func (app *pointerRepositoryBuilder) WithBasePath(basePath []string) pointers.RepositoryBuilder {
	app.basePath = basePath
	return app
}

// Now builds a new pointer repository instance
func (app *pointerRepositoryBuilder) Now() (pointers.Repository, error) {
	if app.basePath != nil && len(app.basePath) <= 0 {
		app.basePath = nil
	}

	if app.basePath == nil {
		return nil, errors.New("the basePath is mandatory in order to build a pointer Repository instance")
	}

	return createPointerRepository(
		app.pointersAdapter,
		app.pointersBuilder,
		app.basePath,
	), nil
}
