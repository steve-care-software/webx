package namespaces

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type namespaceBuilder struct {
	name        string
	description string
	isDeleted   bool
	iterations  delimiters.Delimiter
}

func createNamespaceBuilder() NamespaceBuilder {
	out := namespaceBuilder{
		name:        "",
		description: "",
		isDeleted:   false,
		iterations:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *namespaceBuilder) Create() NamespaceBuilder {
	return createNamespaceBuilder()
}

// WithName adds a name to the builder
func (app *namespaceBuilder) WithName(name string) NamespaceBuilder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *namespaceBuilder) WithDescription(description string) NamespaceBuilder {
	app.description = description
	return app
}

// WithIterations add iterations to the builder
func (app *namespaceBuilder) WithIterations(iterations delimiters.Delimiter) NamespaceBuilder {
	app.iterations = iterations
	return app
}

// IsDeleted flags the builder as deleted
func (app *namespaceBuilder) IsDeleted() NamespaceBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new Namespace instance
func (app *namespaceBuilder) Now() (Namespace, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Namespace instance")
	}

	if app.iterations != nil {
		return createNamespaceWithIterations(app.name, app.description, app.isDeleted, app.iterations), nil
	}

	return createNamespace(app.name, app.description, app.isDeleted), nil
}
