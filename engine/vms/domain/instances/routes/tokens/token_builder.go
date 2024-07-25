package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

type tokenBuilder struct {
	hashAdapter hash.Adapter
	elements    elements.Elements
	cardinality cardinalities.Cardinality
	omission    omissions.Omission
}

func createTokenBuilder(
	hashAdapter hash.Adapter,
) TokenBuilder {
	out := tokenBuilder{
		hashAdapter: hashAdapter,
		elements:    nil,
		cardinality: nil,
		omission:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder(
		app.hashAdapter,
	)
}

// Withelements.Elements add elements to the builder
func (app *tokenBuilder) WithElements(elements elements.Elements) TokenBuilder {
	app.elements = elements
	return app
}

// WithCardinality add cardinality to the builder
func (app *tokenBuilder) WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder {
	app.cardinality = cardinality
	return app
}

// WithOmission add omission to the builder
func (app *tokenBuilder) WithOmission(omission omissions.Omission) TokenBuilder {
	app.omission = omission
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Token instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Token instance")
	}

	data := [][]byte{
		app.elements.Hash().Bytes(),
		app.cardinality.Hash().Bytes(),
	}

	if app.omission != nil {
		data = append(data, app.omission.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.omission != nil {
		return createTokenWithOmission(*pHash, app.elements, app.cardinality, app.omission), nil
	}

	return createToken(*pHash, app.elements, app.cardinality), nil
}
