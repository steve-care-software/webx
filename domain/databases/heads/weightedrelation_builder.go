package heads

import "errors"

type weightedRelationBuilder struct {
	pFrom *uint
	to    WeightedElements
}

func createWeightedRelationBuilder() WeightedRelationBuilder {
	out := weightedRelationBuilder{
		pFrom: nil,
		to:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *weightedRelationBuilder) Create() WeightedRelationBuilder {
	return createWeightedRelationBuilder()
}

// From adds a from index to the builder
func (app *weightedRelationBuilder) From(from uint) WeightedRelationBuilder {
	app.pFrom = &from
	return app
}

// To adds to elements to the builder
func (app *weightedRelationBuilder) To(to WeightedElements) WeightedRelationBuilder {
	app.to = to
	return app
}

// Now builds a new WeightedRelation instance
func (app *weightedRelationBuilder) Now() (WeightedRelation, error) {
	if app.pFrom == nil {
		return nil, errors.New("the from index is mandatory in order to build a WeightedRelation instance")
	}

	if app.to == nil {
		return nil, errors.New("the to elements is mandatory in order to build a WeightedRelation instance")
	}

	return createWeightedRelation(*app.pFrom, app.to), nil
}
