package assignments

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	assignable  assignables.Assignable
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		assignable:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithAssignable adds an assignable to the builder
func (app *builder) WithAssignable(assignable assignables.Assignable) Builder {
	app.assignable = assignable
	return app
}

// Now builds a new Assignment instance
func (app *builder) Now() (Assignment, error) {
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
