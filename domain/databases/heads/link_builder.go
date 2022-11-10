package heads

import "errors"

type linkBuilder struct {
	pFrom *uint
	pTo   *uint
}

func createLinkBuilder() LinkBuilder {
	out := linkBuilder{
		pFrom: nil,
		pTo:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder()
}

// From adds a from index to the builder
func (app *linkBuilder) From(from uint) LinkBuilder {
	app.pFrom = &from
	return app
}

// To adds a to index to the builder
func (app *linkBuilder) To(to uint) LinkBuilder {
	app.pTo = &to
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.pFrom == nil {
		return nil, errors.New("the from index is mandatory in order to buiold a Link instance")
	}

	if app.pTo == nil {
		return nil, errors.New("the to index is mandatory in order to buiold a Link instance")
	}

	return createLink(*app.pFrom, *app.pTo), nil
}
