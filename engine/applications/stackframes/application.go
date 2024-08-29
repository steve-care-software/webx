package stackframes

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/webx/engine/domain/stacks"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

type application struct {
	stackBuilder     stacks.Builder
	frameBuilder     frames.Builder
	variablesBuilder variables.Builder
	variableBuilder  variables.VariableBuilder
	current          stacks.Stack
	path             []string
}

func createApplication(
	stackBuilder stacks.Builder,
	frameBuilder frames.Builder,
	variablesBuilder variables.Builder,
	variableBuilder variables.VariableBuilder,
	current stacks.Stack,
) Application {
	out := application{
		stackBuilder:     stackBuilder,
		frameBuilder:     frameBuilder,
		variablesBuilder: variablesBuilder,
		variableBuilder:  variableBuilder,
		current:          current,
		path:             []string{},
	}

	return &out
}

// Root selects the root stack
func (app *application) Root() {
	app.path = []string{}
}

// Navigate navigates to the stack where the path points
func (app *application) Navigate(path []string) error {
	app.path = path
	return nil
}

// Fetch fetches the stack where our cursor points to
func (app *application) Fetch() (stacks.Stack, error) {
	return app.fetch(app.current, app.path)
}

// Push pushes the stack where our cursor points to
func (app *application) Push(variables variables.Variables) error {
	retParent, err := app.Fetch()
	if err != nil {
		return err
	}

	frame, err := app.frameBuilder.Create().
		WithVariables(variables).
		Now()

	if err != nil {
		return err
	}

	retStack, err := app.stackBuilder.Create().
		WithFrame(frame).
		WithParent(retParent).
		Now()

	if err != nil {
		return err
	}

	retUpdated, err := app.replace(
		retParent,
		app.path,
		retStack,
	)

	if err != nil {
		return err
	}

	app.current = retUpdated
	return nil
}

// Pop pops the stack where our cursor points to
func (app *application) Pop() error {
	retCurrent, err := app.Fetch()
	if err != nil {
		return err
	}

	if retCurrent.HasParent() {
		str := fmt.Sprintf("the current stack (path: %s) could not be popped because it contains no parent stack", strings.Join(app.path, "/"))
		return errors.New(str)
	}

	retParent := retCurrent.Parent()
	retUpdated, err := app.replace(app.current, app.path, retParent)
	if err != nil {
		return err
	}

	app.current = retUpdated
	return nil
}

func (app *application) fetch(current stacks.Stack, path []string) (stacks.Stack, error) {
	if len(path) <= 0 {
		return current, nil
	}

	first := path[0]
	retVariable, err := current.Fetch(first)
	if err != nil {
		return nil, err
	}

	if retVariable.Kind() != variables.KindStack {
		str := fmt.Sprintf("the variable (name: %s) was expected to contain a Stack", first)
		return nil, errors.New(str)
	}

	value := retVariable.Value()
	if casted, ok := value.(stacks.Stack); ok {
		return app.fetch(casted, path[1:])
	}

	str := fmt.Sprintf("the variable (name: %s) was expected to contain a Stack, the kind was correctly set but the value could not be casted properly", first)
	return nil, errors.New(str)
}

func (app *application) replace(
	current stacks.Stack,
	path []string,
	replacement stacks.Stack,
) (stacks.Stack, error) {
	if len(path) <= 1 {
		variableName := path[0]
		return app.repalceStackInVariable(current, variableName, replacement)
	}

	first := path[0]
	retVariable, err := current.Fetch(first)
	if err != nil {
		return nil, err
	}

	if retVariable.Kind() != variables.KindStack {
		str := fmt.Sprintf("the variable (name: %s) was expected to contain a Stack", first)
		return nil, errors.New(str)
	}

	value := retVariable.Value()
	if casted, ok := value.(stacks.Stack); ok {
		retUpdatedStack, err := app.fetch(casted, path[1:])
		if err != nil {
			return nil, err
		}

		return app.repalceStackInVariable(current, first, retUpdatedStack)
	}

	str := fmt.Sprintf("the variable (name: %s) was expected to contain a Stack, the kind was correctly set but the value could not be casted properly", first)
	return nil, errors.New(str)
}

func (app *application) repalceStackInVariable(
	current stacks.Stack,
	variableName string,
	replacement stacks.Stack,
) (stacks.Stack, error) {
	frame := current.Frame()
	if !frame.HasVariables() {
		str := fmt.Sprintf("the variable (name: %s) could not be retrieved because the frame contains no variable", variableName)
		return nil, errors.New(str)
	}

	variablesList := frame.Variables().List()
	for idx, oneVariable := range variablesList {
		name := oneVariable.Name()
		if name == variableName {
			kind := oneVariable.Kind()
			replaceIfExists := oneVariable.ReplaceIfExists()

			builder := app.variableBuilder.Create().
				WithName(name).
				WithKind(kind).
				WithValue(replacement)

			if replaceIfExists {
				builder.ReplaceIfExists()
			}

			ins, err := builder.Now()
			if err != nil {
				return nil, err
			}

			variablesList[idx] = ins
			continue
		}
	}

	variables, err := app.variablesBuilder.Create().
		WithList(variablesList).
		Now()

	if err != nil {
		return nil, err
	}

	frameIns, err := app.frameBuilder.Create().
		WithVariables(variables).
		Now()

	if err != nil {
		return nil, err
	}

	builder := app.stackBuilder.Create().WithFrame(frameIns)
	if current.HasParent() {
		parent := current.Parent()
		builder.WithParent(parent)
	}

	return builder.Now()
}
