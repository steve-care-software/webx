package entities

import "errors"

type identifiersBuilder struct {
	list []Identifier
}

func createIdentifiersBuilder() IdentifiersBuilder {
	out := identifiersBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *identifiersBuilder) Create() IdentifiersBuilder {
	return createIdentifiersBuilder()
}

// WithList adds a list to the builder
func (app *identifiersBuilder) WithList(list []Identifier) IdentifiersBuilder {
	app.list = list
	return app
}

// Now builds a new Identifiers instance
func (app *identifiersBuilder) Now() (Identifiers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Identifier in order to build an Identifiers instance")
	}

	return createIdentifiers(app.list), nil
}
