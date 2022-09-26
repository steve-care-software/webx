package grammars

import "errors"

type suitesBuilder struct {
	list []Suite
}

func createSuitesBuilder() SuitesBuilder {
	out := suitesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suitesBuilder) Create() SuitesBuilder {
	return createSuitesBuilder()
}

// WithList adds a list to the builder
func (app *suitesBuilder) WithList(list []Suite) SuitesBuilder {
	app.list = list
	return app
}

// Now builds a new Suites instance
func (app *suitesBuilder) Now() (Suites, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Suite in order to build a Suites instance")
	}

	return createSuites(app.list), nil
}
