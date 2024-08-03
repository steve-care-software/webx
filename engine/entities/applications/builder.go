package applications

import (
	"errors"

	"github.com/steve-care-software/webx/engine/entities/domain/entities"
	hash_applications "github.com/steve-care-software/webx/engine/hashes/applications"
)

type builder struct {
	entityAdapter entities.Adapter
	hashApp       hash_applications.Application
}

func createBuilder(
	entityAdapter entities.Adapter,
) Builder {
	out := builder{
		entityAdapter: entityAdapter,
		hashApp:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.entityAdapter,
	)
}

// WithHash adds an hash application to the builder
func (app *builder) WithHash(hashApp hash_applications.Application) Builder {
	app.hashApp = hashApp
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.hashApp == nil {
		return nil, errors.New("the hash application is mandatory in order to build an Application")
	}

	return createApplication(
		app.hashApp,
		app.entityAdapter,
	), nil
}
