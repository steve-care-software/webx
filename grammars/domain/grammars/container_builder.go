package grammars

import "errors"

type containerBuilder struct {
	element Element
	compose Compose
}

func createContainerBuilder() ContainerBuilder {
	out := containerBuilder{
		element: nil,
		compose: nil,
	}

	return &out
}

// Create initializes the builder
func (app *containerBuilder) Create() ContainerBuilder {
	return createContainerBuilder()
}

// WithElement adds an element to the builder
func (app *containerBuilder) WithElement(element Element) ContainerBuilder {
	app.element = element
	return app
}

// WithCompose adds a compose to the builder
func (app *containerBuilder) WithCompose(compose Compose) ContainerBuilder {
	app.compose = compose
	return app
}

// Now builds a new Container instance
func (app *containerBuilder) Now() (Container, error) {
	if app.element != nil {
		return createContainerWithElement(app.element), nil
	}

	if app.compose != nil {
		return createContainerWithCompose(app.compose), nil
	}

	return nil, errors.New("the Container is invalid")
}
