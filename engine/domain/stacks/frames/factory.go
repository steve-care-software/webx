package frames

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

// Create creates a new frame
func (app *factory) Create() (Frame, error) {
	return app.builder.Now()
}
