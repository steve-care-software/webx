package stacks

type factory struct {
	framesBuilder FramesBuilder
	frameBuilder  FrameBuilder
}

func createFactory(
	framesBuilder FramesBuilder,
	frameBuilder FrameBuilder,
) Factory {
	out := factory{
		framesBuilder: framesBuilder,
		frameBuilder:  frameBuilder,
	}

	return &out
}

// Create creates a new stack
func (app *factory) Create() (Stack, error) {
	head, err := app.frameBuilder.Create().Now()
	if err != nil {
		return nil, err
	}

	frames, err := app.framesBuilder.Create().WithList([]Frame{
		head,
	}).Now()

	if err != nil {
		return nil, err
	}

	return createStackWithSingleFrame(frames, head), nil
}
