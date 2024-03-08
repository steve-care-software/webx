package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
)

type linkBuilder struct {
	hashAdapter hash.Adapter
	origin      origins.Origin
	elements    elements.Elements
}

func createLinkBuilder(
	hashAdapter hash.Adapter,
) LinkBuilder {
	out := linkBuilder{
		hashAdapter: hashAdapter,
		origin:      nil,
		elements:    nil,
	}

	return &out
}

// LinkBuilder initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder(
		app.hashAdapter,
	)
}

// WithOrigin adds an origin to the builder
func (app *linkBuilder) WithOrigin(origin origins.Origin) LinkBuilder {
	app.origin = origin
	return app
}

// WithElements add elements to the builder
func (app *linkBuilder) WithElements(elements elements.Elements) LinkBuilder {
	app.elements = elements
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.origin == nil {
		return nil, errors.New("the origin is mandatory in order to build a Link instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Link instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.origin.Hash().Bytes(),
		app.elements.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLink(*pHash, app.origin, app.elements), nil
}
