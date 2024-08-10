package profiles

type builder struct {
	name        string
	description string
	namespaces  []string
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		description: "",
		namespaces:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithNamespaces add namespaces to the builder
func (app *builder) WithNamespaces(namespaces []string) Builder {
	app.namespaces = namespaces
	return app
}

// Now builds a new Profile instance
func (app *builder) Now() (Profile, error) {
	if app.namespaces != nil && len(app.namespaces) <= 0 {
		app.namespaces = nil
	}

	if app.namespaces != nil {
		return createProfileWithNamespaces(app.name, app.description, app.namespaces), nil
	}

	return createProfile(app.name, app.description), nil
}
