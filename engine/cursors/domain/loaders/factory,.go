package loaders

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

// Create creates a new loader instance
func (app *factory) Create() (Loader, error) {
	return app.builder.Create().Now()
}
