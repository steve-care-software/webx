package selectors

import "errors"

type elementBuilder struct {
	name   string
	pIndex *uint
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		name:   "",
		pIndex: nil,
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

// WithIndex adds an index to the builder
func (app *elementBuilder) WithIndex(index uint) ElementBuilder {
	app.pIndex = &index
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Element instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Element instance")
	}

	return createElement(app.name, *app.pIndex), nil
}
