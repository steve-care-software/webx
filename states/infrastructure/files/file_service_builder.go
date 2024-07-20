package files

import (
	"errors"

	"github.com/steve-care-software/datastencil/states/domain/files"
)

type fileServiceBuilder struct {
	repositoryBuilder files.RepositoryBuilder
	innerPath         []string
	basePath          []string
}

func createFileServiceBuilder(
	repositoryBuilder files.RepositoryBuilder,
	innerPath []string,
) files.ServiceBuilder {
	out := fileServiceBuilder{
		repositoryBuilder: repositoryBuilder,
		innerPath:         innerPath,
		basePath:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *fileServiceBuilder) Create() files.ServiceBuilder {
	return createFileServiceBuilder(
		app.repositoryBuilder,
		app.innerPath,
	)
}

// WithBasePath adds a basePath to the builder
func (app *fileServiceBuilder) WithBasePath(basePath []string) files.ServiceBuilder {
	app.basePath = basePath
	return app
}

// Now builds a new Service instance
func (app *fileServiceBuilder) Now() (files.Service, error) {
	if app.basePath != nil && len(app.basePath) <= 0 {
		app.basePath = []string{}
	}

	if app.innerPath != nil && len(app.innerPath) <= 0 {
		app.innerPath = []string{}
	}

	repository, err := app.repositoryBuilder.Create().WithBasePath(app.basePath).Now()
	if err != nil {
		return nil, err
	}

	basePath := append(app.basePath, app.innerPath...)
	if len(basePath) <= 0 {
		return nil, errors.New("the combined basePath and innerPath must not be empty")
	}

	return createFileService(
		repository,
		basePath,
	), nil
}
