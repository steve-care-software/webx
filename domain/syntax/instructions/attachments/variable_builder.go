package attachments

import "errors"

type variableBuilder struct {
	current string
	target  string
}

func createVariableBuilder() VariableBuilder {
	out := variableBuilder{
		current: "",
		target:  "",
	}

	return &out
}

// Create initializes the builder
func (app *variableBuilder) Create() VariableBuilder {
	return createVariableBuilder()
}

// WithCurrent adds a current variable to the builder
func (app *variableBuilder) WithCurrent(current string) VariableBuilder {
	app.current = current
	return app
}

// WithTarget adds a target variable to the builder
func (app *variableBuilder) WithTarget(target string) VariableBuilder {
	app.target = target
	return app
}

// Now builds a new Variable instance
func (app *variableBuilder) Now() (Variable, error) {
	if app.current == "" {
		return nil, errors.New("the current variable is mandatory in order to build a Variable instance")
	}

	if app.target == "" {
		return nil, errors.New("the target variable is mandatory in order to build a Variable instance")
	}

	return createVariable(app.current, app.target), nil
}
