package files

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

type layerRepositoryBuilder struct {
	//pointerRepositoryBuilder pointers.RepositoryBuilder
	adapter  layers.Adapter
	basePath []string
}

func createLayerRepositoryBuilder(
	//pointerRepositoryBuilder pointers.RepositoryBuilder,
	adapter layers.Adapter,
) layers.RepositoryBuilder {
	out := layerRepositoryBuilder{
		//pointerRepositoryBuilder: pointerRepositoryBuilder,
		adapter:  adapter,
		basePath: nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerRepositoryBuilder) Create() layers.RepositoryBuilder {
	return createLayerRepositoryBuilder(
		//app.pointerRepositoryBuilder,
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

	/*pointerRepository, err := app.pointerRepositoryBuilder.Create().
		WithBasePath(app.basePath).
		Now()

	if err != nil {
		return nil, err
	}*/

	return createLayerRepository(
		//pointerRepository,
		app.adapter,
	), nil
}
