package stacks

import "github.com/steve-care-software/webx/engine/domain/stacks/frames"

type factory struct {
	builder      Builder
	frameFactory frames.Factory
}

func createFactory(
	builder Builder,
	frameFactory frames.Factory,
) Factory {
	out := factory{
		builder:      builder,
		frameFactory: frameFactory,
	}

	return &out
}

// Create creates a new stack
func (app *factory) Create() (Stack, error) {
	frame, err := app.frameFactory.Create()
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithFrame(frame).
		Now()
}
