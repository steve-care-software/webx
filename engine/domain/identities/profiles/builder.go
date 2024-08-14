package profiles

type builder struct {
	name        string
	description string
	packages    []string
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		description: "",
		packages:    nil,
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

// WithPackages add packages to the builder
func (app *builder) WithPackages(packages []string) Builder {
	app.packages = packages
	return app
}

// Now builds a new Profile instance
func (app *builder) Now() (Profile, error) {
	if app.packages != nil && len(app.packages) <= 0 {
		app.packages = nil
	}

	if app.packages != nil {
		return createProfileWithPackages(app.name, app.description, app.packages), nil
	}

	return createProfile(app.name, app.description), nil
}
