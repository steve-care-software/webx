package suites

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type suiteBuilder struct {
	name    string
	element elements.Element
	isFail  bool
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		name:    "",
		element: nil,
		isFail:  false,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder()
}

// WithName adds a name to the builder
func (app *suiteBuilder) WithName(name string) SuiteBuilder {
	app.name = name
	return app
}

// WithElement adds an element to the builder
func (app *suiteBuilder) WithElement(element elements.Element) SuiteBuilder {
	app.element = element
	return app
}

// IsFail flags the suite as fail
func (app *suiteBuilder) IsFail() SuiteBuilder {
	app.isFail = true
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Suite instance")
	}

	return createSuite(app.name, app.element, app.isFail), nil
}
