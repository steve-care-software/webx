package references

type contentFactory struct {
	builder ContentBuilder
}

func createContentFactory(
	builder ContentBuilder,
) ContentFactory {
	out := contentFactory{
		builder: builder,
	}

	return &out
}

// Create creates a new content instance
func (app *contentFactory) Create() (Content, error) {
	return app.builder.Create().Now()
}
