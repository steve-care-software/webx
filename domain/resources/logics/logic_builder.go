package logics

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/resources/logics/bridges"
	"github.com/steve-care-software/datastencil/domain/resources/logics/references"
)

type logicBuilder struct {
	hashAdapter hash.Adapter
	link        links.Link
	bridges     bridges.Bridges
	references  references.References
}

func createLogicBuilder(
	hashAdapter hash.Adapter,
) LogicBuilder {
	out := logicBuilder{
		hashAdapter: hashAdapter,
		link:        nil,
		bridges:     nil,
		references:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *logicBuilder) Create() LogicBuilder {
	return createLogicBuilder(
		app.hashAdapter,
	)
}

// WithLink adds a link to the builder
func (app *logicBuilder) WithLink(link links.Link) LogicBuilder {
	app.link = link
	return app
}

// WithBridges add bridges to the builder
func (app *logicBuilder) WithBridges(bridges bridges.Bridges) LogicBuilder {
	app.bridges = bridges
	return app
}

// WithReferences add references to the builder
func (app *logicBuilder) WithReferences(references references.References) LogicBuilder {
	app.references = references
	return app
}

// Now builds a new Logic instance
func (app *logicBuilder) Now() (Logic, error) {
	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Logic instance")
	}

	if app.bridges == nil {
		return nil, errors.New("the bridges is mandatory in order to build a Logic instance")
	}

	data := [][]byte{
		app.link.Hash().Bytes(),
		app.bridges.Hash().Bytes(),
	}

	if app.references != nil {
		data = append(data, app.references.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.references != nil {
		return createLogicWithReferences(*pHash, app.link, app.bridges, app.references), nil
	}

	return createLogic(*pHash, app.link, app.bridges), nil
}
