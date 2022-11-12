package references

import "errors"

type linksBuilder struct {
	list []Link
}

func createLinksBuilder() LinksBuilder {
	out := linksBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linksBuilder) Create() LinksBuilder {
	return createLinksBuilder()
}

// WithList adds a list to the builder
func (app *linksBuilder) WithList(list []Link) LinksBuilder {
	app.list = list
	return app
}

// Now builds a new Links instance
func (app *linksBuilder) Now() (Links, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Link in order to build an Links instance")
	}

	return createLinks(app.list), nil
}
