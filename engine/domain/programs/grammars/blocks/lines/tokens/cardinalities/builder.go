package cardinalities

type builder struct {
	min  uint
	pMax *uint
}

func createBuilder() Builder {
	out := builder{
		min:  0,
		pMax: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMin adds a min to the builder
func (app *builder) WithMin(min uint) Builder {
	app.min = min
	return app
}

// WithMax adds a max to the builder
func (app *builder) WithMax(max uint) Builder {
	app.pMax = &max
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	if app.pMax != nil {
		return createCardinalityWithMax(app.min, app.pMax), nil
	}

	return createCardinality(app.min), nil
}
