package stacks

import "errors"

type builder struct {
	framesBuilder FramesBuilder
	frames        Frames
}

func createBuilder(
	framesBuilder FramesBuilder,
) Builder {
	out := builder{
		framesBuilder: framesBuilder,
		frames:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.framesBuilder,
	)
}

// WithFrames add frames to the builder
func (app *builder) WithFrames(frames Frames) Builder {
	app.frames = frames
	return app
}

// Now builds a new Stack instance
func (app *builder) Now() (Stack, error) {
	if app.frames == nil {
		return nil, errors.New("the frames is mandatory in order to build a stack instance")
	}

	list := app.frames.List()
	amount := len(list)
	if amount <= 1 {
		return createStackWithSingleFrame(app.frames, list[0]), nil
	}

	body, err := app.framesBuilder.Create().
		WithList(list[:amount-2]).
		Now()

	if err != nil {
		return nil, err
	}

	return createStackWithFrames(app.frames, body, list[amount-1]), nil
}
