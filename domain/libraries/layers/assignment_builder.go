package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type assignmentBuilder struct {
	hashAdapter hash.Adapter
	name        string
	assignable  Assignable
}

func createAssignmentBuilder(
	hashAdapter hash.Adapter,
) AssignmentBuilder {
	out := assignmentBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		assignable:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
	return createAssignmentBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *assignmentBuilder) WithName(name string) AssignmentBuilder {
	app.name = name
	return app
}

// WithAssignable adds an assignable to the builder
func (app *assignmentBuilder) WithAssignable(assignable Assignable) AssignmentBuilder {
	app.assignable = assignable
	return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Assignment instance")
	}

	if app.assignable == nil {
		return nil, errors.New("the assignable is mandatory in order to build an Assignment instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.assignable.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createAssignment(*pHash, app.name, app.assignable), nil
}
