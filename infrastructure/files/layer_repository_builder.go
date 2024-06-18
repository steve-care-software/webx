package files

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
)

type layerRepositoryBuilder struct {
	adapter  layers.Adapter
	basePath []string
}

func createLayerRepositoryBuilder(
	adapter layers.Adapter,
) layers.RepositoryBuilder {
	out := layerRepositoryBuilder{
		adapter:  adapter,
		basePath: nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerRepositoryBuilder) Create() layers.RepositoryBuilder {
	return createLayerRepositoryBuilder(
		app.adapter,
	)
}

// WithBasePath adds a base path to the builder
func (app *layerRepositoryBuilder) WithBasePath(basePath []string) layers.RepositoryBuilder {
	app.basePath = basePath
	return app
}

// Now builds a new layer repository instance
func (app *layerRepositoryBuilder) Now() (layers.Repository, error) {
	if app.basePath != nil && len(app.basePath) <= 0 {
		app.basePath = nil
	}

	if app.basePath == nil {
		return nil, errors.New("the basePath is mandatory in order to build a layer Repository instance")
	}

	return createFileRepository(
		app.adapter,
		app.basePath,
	), nil
}
