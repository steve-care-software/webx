package coverages

import "errors"

type elementBuilder struct {
	name  string
	value []byte
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		name:  "",
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithName adds a name to the builder
func (app *elementBuilder) WithName(name string) ElementBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *elementBuilder) WithValue(value []byte) ElementBuilder {
	app.value = value
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Element instance")
	}

	if app.value != nil {
		return createElementWithValue(app.name, app.value), nil
	}

	return createElement(app.name), nil
}
