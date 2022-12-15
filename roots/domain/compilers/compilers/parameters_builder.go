package compilers

import "errors"

type parametersBuilder struct {
	list []Parameter
}

func createParametersBuilder() ParametersBuilder {
	out := parametersBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *parametersBuilder) Create() ParametersBuilder {
	return createParametersBuilder()
}

// WithList add parameters to the builder
func (app *parametersBuilder) WithList(list []Parameter) ParametersBuilder {
	app.list = list
	return app
}

// Now builds a new Parameters instance
func (app *parametersBuilder) Now() (Parameters, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Parameter in order to build a Parameters instance")
	}

	return createParameters(app.list), nil
}
