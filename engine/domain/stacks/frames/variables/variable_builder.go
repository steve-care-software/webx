package variables

import (
	"errors"
	"fmt"
)

type variableBuilder struct {
	name            string
	value           any
	pKind           *uint8
	replaceIfExists bool
}

func createVariableBuilder() VariableBuilder {
	out := variableBuilder{
		name:            "",
		value:           nil,
		pKind:           nil,
		replaceIfExists: false,
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
func (app *variableBuilder) WithValue(value any) VariableBuilder {
	app.value = value
	return app
}

// WithKind adds a kind to the builder
func (app *variableBuilder) WithKind(kind uint8) VariableBuilder {
	app.pKind = &kind
	return app
}

// ReplaceIfExists flags the builder as replaceIfExists
func (app *variableBuilder) ReplaceIfExists() VariableBuilder {
	app.replaceIfExists = true
	return app
}

// Now builds a new Variable instance
func (app *variableBuilder) Now() (Variable, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Variable instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Variable instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Variable instance")
	}

	kind := *app.pKind
	if kind > KindStack {
		str := fmt.Sprintf("the kind (%d) is invalid", kind)
		return nil, errors.New(str)
	}

	return createVariable(
		app.name,
		app.value,
		kind,
		app.replaceIfExists,
	), nil
}
