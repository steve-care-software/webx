package references

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

// Create generates a new Reference instance
func (app *factory) Create() (Reference, error) {
	return app.builder.Create().Now()
}
