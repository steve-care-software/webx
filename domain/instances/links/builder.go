package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/references"
)

type builder struct {
	hashAdapter hash.Adapter
	elements    elements.Elements
	references  references.References
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		elements:    nil,
		references:  nil,
	}

	return &out
}

// Builder initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithElements add elements to the builder
func (app *builder) WithElements(elements elements.Elements) Builder {
	app.elements = elements
	return app
}

// WithReferences add references to the builder
func (app *builder) WithReferences(references references.References) Builder {
	app.references = references
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
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
