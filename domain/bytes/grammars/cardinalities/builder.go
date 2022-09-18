package cardinalities

import "errors"

type builder struct {
	pMin *uint
	pMax *uint
}

func createBuilder() Builder {
	out := builder{
		pMin: nil,
		pMax: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMin adds a minimum to the builder
func (app *builder) WithMin(min uint) Builder {
	app.pMin = &min
	return app
}

// WithMax adds a maximum to the builder
func (app *builder) WithMax(max uint) Builder {
	app.pMax = &max
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	if app.pMin == nil {
		return nil, errors.New("the minimum is mandatory in order to build a Cardinality instance")
	}

	if app.pMax != nil {
		return createCardinalityWithMax(*app.pMin, app.pMax), nil
	}

	return createCardinality(*app.pMin), nil
}
