package publics

import "errors"

type builder struct {
	publics []Public
}

func createBuilder() Builder {
	out := builder{
		publics: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add publics to the builder
func (app *builder) WithList(publics []Public) Builder {
	app.publics = publics
	return app
}

// Now builds a new Publics instance
func (app *builder) Now() (Publics, error) {
	if app.publics != nil && len(app.publics) <= 0 {
		app.publics = nil
	}

	if app.publics == nil {
		return nil, errors.New("there must be at least 1 Public in order to build an Publics instance")
	}

	return createPublics(app.publics), nil
}
