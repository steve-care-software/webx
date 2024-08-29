package stacks

type factory struct {
	builder Builder
}

func createFactory(
	builder Builder,
) Factory {
	out := factory{
		builder: builder,
	}

	return &out
}

// Create creates a new stack
func (app *factory) Create() (Stack, error) {
	return app.builder.Now()
}
