package stacks

import (
	"errors"
)

type framesBuilder struct {
	list []Frame
}

func createFramesBuilder() FramesBuilder {
	out := framesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *framesBuilder) Create() FramesBuilder {
	return createFramesBuilder()
}

// WithList adds a list to the builder
func (app *framesBuilder) WithList(list []Frame) FramesBuilder {
	app.list = list
	return app
}

// Now builds a new Frames instance
func (app *framesBuilder) Now() (Frames, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Frame in order to build an Frames instance")
	}

	return createFrames(app.list), nil
}
