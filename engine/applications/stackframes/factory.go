package stackframes

import (
	"github.com/steve-care-software/webx/engine/domain/stacks"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

type factory struct {
	stackFactory     stacks.Factory
	stackBuilder     stacks.Builder
	frameBuilder     frames.Builder
	variablesBuilder variables.Builder
	variableBuilder  variables.VariableBuilder
}

func cretaeFactory(
	stackFactory stacks.Factory,
	stackBuilder stacks.Builder,
	frameBuilder frames.Builder,
	variablesBuilder variables.Builder,
	variableBuilder variables.VariableBuilder,
) Factory {
	out := factory{
		stackFactory:     stackFactory,
		stackBuilder:     stackBuilder,
		frameBuilder:     frameBuilder,
		variablesBuilder: variablesBuilder,
		variableBuilder:  variableBuilder,
	}

	return &out
}

// Create creates the application
func (app *factory) Create() (Application, error) {
	current, err := app.stackFactory.Create()
	if err != nil {
		return nil, err
	}

	return createApplication(
		app.stackBuilder,
		app.frameBuilder,
		app.variablesBuilder,
		app.variableBuilder,
		current,
	), nil
}
