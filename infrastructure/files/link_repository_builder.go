package files

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
)

type linkRepositoryBuilder struct {
	adapter  links.Adapter
	basePath []string
}

func createLinkRepositoryBuilder(
	adapter links.Adapter,
) links.RepositoryBuilder {
	out := linkRepositoryBuilder{
		adapter:  adapter,
		basePath: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkRepositoryBuilder) Create() links.RepositoryBuilder {
	return createLinkRepositoryBuilder(
		app.adapter,
	)
}

// WithBasePath adds a base path to the builder
func (app *linkRepositoryBuilder) WithBasePath(basePath []string) links.RepositoryBuilder {
	app.basePath = basePath
	return app
}

// Now builds a new Repository instance
func (app *linkRepositoryBuilder) Now() (links.Repository, error) {
	if app.basePath != nil && len(app.basePath) <= 0 {
		app.basePath = nil
	}

	if app.basePath == nil {
		return nil, errors.New("the basePath is mandatory in order to build a link Repository instance")
	}

	return createLinkRepository(
		app.adapter,
		app.basePath,
	), nil
}
