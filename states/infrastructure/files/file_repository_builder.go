package files

import (
	"errors"

	"github.com/steve-care-software/datastencil/states/domain/files"
)

type fileRepositoryBuilder struct {
	innerPath []string
	basePath  []string
}

func createFileRepositoryBuilder(
	innerPath []string,
) files.RepositoryBuilder {
	out := fileRepositoryBuilder{
		innerPath: innerPath,
		basePath:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *fileRepositoryBuilder) Create() files.RepositoryBuilder {
	return createFileRepositoryBuilder(
		app.innerPath,
	)
}

// WithBasePath adds a basePath to the builder
func (app *fileRepositoryBuilder) WithBasePath(basePath []string) files.RepositoryBuilder {
	app.basePath = basePath
	return app
}

// Now builds a new Repository instance
func (app *fileRepositoryBuilder) Now() (files.Repository, error) {
	if app.basePath != nil && len(app.basePath) <= 0 {
		app.basePath = []string{}
	}

	if app.innerPath != nil && len(app.innerPath) <= 0 {
		app.innerPath = []string{}
	}

	basePath := append(app.basePath, app.innerPath...)
	if len(basePath) <= 0 {
		return nil, errors.New("the combined basePath and innerPath must not be empty")
	}

	return createFileRepository(
		basePath,
	), nil
}
