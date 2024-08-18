package replacements

import (
	"errors"
)

type builder struct {
	list []Replacement
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Replacement) Builder {
	app.list = list
	return app
}

// Now builds a new Replacements instance
func (app *builder) Now() (Replacements, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Replacement in order to build a Replacements instance")
	}

	return createReplacements(app.list), nil
}
