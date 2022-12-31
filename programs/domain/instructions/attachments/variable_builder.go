package attachments

import "errors"

type variableBuilder struct {
	current []byte
	pTarget *uint
}

func createVariableBuilder() VariableBuilder {
	out := variableBuilder{
		current: nil,
		pTarget: nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableBuilder) Create() VariableBuilder {
	return createVariableBuilder()
}

// WithCurrent adds a current variable to the builder
func (app *variableBuilder) WithCurrent(current []byte) VariableBuilder {
	app.current = current
	return app
}

// WithTarget adds a target variable to the builder
func (app *variableBuilder) WithTarget(target uint) VariableBuilder {
	app.pTarget = &target
	return app
}

// Now builds a new Variable instance
func (app *variableBuilder) Now() (Variable, error) {
	if app.current == nil {
		return nil, errors.New("the current variable is mandatory in order to build a Variable instance")
	}

	if app.pTarget == nil {
		return nil, errors.New("the target variable is mandatory in order to build a Variable instance")
	}

	return createVariable(app.current, *app.pTarget), nil
}
