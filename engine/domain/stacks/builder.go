package stacks

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
)

type builder struct {
	frame  frames.Frame
	parent Stack
}

func createBuilder() Builder {
	out := builder{
		frame:  nil,
		parent: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithFrame adds a frame to the builder
func (app *builder) WithFrame(frame frames.Frame) Builder {
	app.frame = frame
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent Stack) Builder {
	app.parent = parent
	return app
}

// Now builds a new Stack instance
func (app *builder) Now() (Stack, error) {
	if app.frame == nil {
		return nil, errors.New("the frame is mandatory in order to build a Stack instance")
	}

	if app.parent != nil {
		return createStackWithParent(
			app.frame,
			app.parent,
		), nil
	}

	return createStack(
		app.frame,
	), nil
}
