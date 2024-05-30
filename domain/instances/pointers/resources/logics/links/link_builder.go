package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/references"
)

type linkBuilder struct {
	hashAdapter hash.Adapter
	elements    elements.Elements
	references  references.References
}

func createLinkBuilder(
	hashAdapter hash.Adapter,
) LinkBuilder {
	out := linkBuilder{
		hashAdapter: hashAdapter,
		elements:    nil,
		references:  nil,
	}

	return &out
}

// LinkBuilder initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder(
		app.hashAdapter,
	)
}

// WithElements add elements to the builder
func (app *linkBuilder) WithElements(elements elements.Elements) LinkBuilder {
	app.elements = elements
	return app
}

// WithReferences add references to the builder
func (app *linkBuilder) WithReferences(references references.References) LinkBuilder {
	app.references = references
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Link instance")
	}

	data := [][]byte{
		app.elements.Hash().Bytes(),
	}

	if app.references != nil {
		data = append(data, app.references.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.references != nil {
		return createLinkWithReferences(*pHash, app.elements, app.references), nil
	}

	return createLink(*pHash, app.elements), nil
}
