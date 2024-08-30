package resources

import "errors"

type resourceBuilder struct {
	name         string
	relativePath string
}

func createResourceBuilder() ResourceBuilder {
	out := resourceBuilder{
		name:         "",
		relativePath: "",
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder()
}

// WithName adds a name to the builder
func (app *resourceBuilder) WithName(name string) ResourceBuilder {
	app.name = name
	return app
}

// WithRelativePath adds a relativePath to the builder
func (app *resourceBuilder) WithRelativePath(relativePath string) ResourceBuilder {
	app.relativePath = relativePath
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Resource instance")
	}

	if app.relativePath == "" {
		return nil, errors.New("the relativePath is mandatory in order to build a Resource instance")
	}

	return createResource(
		app.name,
		app.relativePath,
	), nil
}
