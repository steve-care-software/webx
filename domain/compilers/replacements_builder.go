package compilers

import "errors"

type replacementsBuilder struct {
	list []Replacement
}

func createReplacementsBuilder() ReplacementsBuilder {
	out := replacementsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *replacementsBuilder) Create() ReplacementsBuilder {
	return createReplacementsBuilder()
}

// WithList adds a list to the builder
func (app *replacementsBuilder) WithList(list []Replacement) ReplacementsBuilder {
	app.list = list
	return app
}

// Now builds a new Replacements instance
func (app *replacementsBuilder) Now() (Replacements, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Replacement in order to build a Replacements instance")
	}

	return createReplacements(app.list), nil
}
