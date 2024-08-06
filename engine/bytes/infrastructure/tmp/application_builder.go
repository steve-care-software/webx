package tmp

import (
	"errors"

	"github.com/steve-care-software/webx/engine/bytes/applications"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type applicationBuilder struct {
	namespaceAdapter  namespaces.Adapter
	namespacesBuilder namespaces.Builder
	namespaceBuilder  namespaces.NamespaceBuilder
	hashAdapter       hash.Adapter
	basepath          []string
}

func createApplicationBuilder(
	namespaceAdapter namespaces.Adapter,
	namespacesBuilder namespaces.Builder,
	namespaceBuilder namespaces.NamespaceBuilder,
	hashAdapter hash.Adapter,
) applications.Builder {
	out := applicationBuilder{
		namespaceAdapter:  namespaceAdapter,
		namespacesBuilder: namespacesBuilder,
		namespaceBuilder:  namespaceBuilder,
		hashAdapter:       hashAdapter,
		basepath:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() applications.Builder {
	return createApplicationBuilder(
		app.namespaceAdapter,
		app.namespacesBuilder,
		app.namespaceBuilder,
		app.hashAdapter,
	)
}

// WithBasePath adds a basePath to the builder
func (app *applicationBuilder) WithBasePath(basepath []string) applications.Builder {
	app.basepath = basepath
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (applications.Application, error) {
	if app.basepath != nil && len(app.basepath) <= 0 {
		app.basepath = nil
	}

	if app.basepath == nil {
		return nil, errors.New("the basePath is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.namespaceAdapter,
		app.namespacesBuilder,
		app.namespaceBuilder,
		app.hashAdapter,
		app.basepath,
	), nil
}
