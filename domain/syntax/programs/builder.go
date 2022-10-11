package programs

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/programs/applications"
)

type builder struct {
	assignments []applications.Assignment
	outputs     []string
}

func createBuilder() Builder {
	out := builder{
		assignments: nil,
		outputs:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithAssignments add assignments to the builder
func (app *builder) WithAssignments(assignments []applications.Assignment) Builder {
	app.assignments = assignments
	return app
}

// WithOutputs add outputs to the builder
func (app *builder) WithOutputs(outputs []string) Builder {
	app.outputs = outputs
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.assignments != nil && len(app.assignments) <= 0 {
		app.assignments = nil
	}

	if app.assignments == nil {
		return nil, errors.New("the assignments are mandatory in order to build a Program instance")
	}

	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	if app.outputs != nil {
		return createProgramWithOutputs(app.assignments, app.outputs), nil
	}

	return createProgram(app.assignments), nil
}
