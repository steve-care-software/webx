package outputs

import "errors"

type variableBuilder struct {
	name  string
	value interface{}
}

func createVariableBuilder() VariableBuilder {
	out := variableBuilder{
		name:  "",
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableBuilder) Create() VariableBuilder {
	return createVariableBuilder()
}

// WithName adds a name to the builder
func (app *variableBuilder) WithName(name string) VariableBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *variableBuilder) WithValue(value interface{}) VariableBuilder {
	app.value = value
	return app
}

// Now builds a new Variable instance
func (app *variableBuilder) Now() (Variable, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Variable instance")
	}

	if app.value != nil {
		return createVariableWithValue(app.name, app.value), nil
	}

	return createVariable(app.name), nil
}
