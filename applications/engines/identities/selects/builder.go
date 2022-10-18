package selects

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/identities"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

type builder struct {
	builder              identities.Builder
	modificationsBuilder modifications.Builder
	repository           identities.Repository
	service              identities.Service
	name                 string
}

func createBuilder(
	identityBuilder identities.Builder,
	modificationsBuilder modifications.Builder,
	repository identities.Repository,
	service identities.Service,
) Builder {
	out := builder{
		builder:              identityBuilder,
		modificationsBuilder: modificationsBuilder,
		repository:           repository,
		service:              service,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.builder,
		app.modificationsBuilder,
		app.repository,
		app.service,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.builder,
		app.modificationsBuilder,
		app.repository,
		app.service,
		app.name,
	), nil
}
