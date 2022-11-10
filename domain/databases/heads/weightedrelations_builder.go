package heads

import "errors"

type weightedRelationsBuilder struct {
	list []WeightedRelation
}

func createWeightedRelationsBuilder() WeightedRelationsBuilder {
	out := weightedRelationsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *weightedRelationsBuilder) Create() WeightedRelationsBuilder {
	return createWeightedRelationsBuilder()
}

// WithList adds a list to the builder
func (app *weightedRelationsBuilder) WithList(list []WeightedRelation) WeightedRelationsBuilder {
	app.list = list
	return app
}

// Now builds a new WeightedRelations instance
func (app *weightedRelationsBuilder) Now() (WeightedRelations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 WeightedRelation in order to build an WeightedRelations instance")
	}

	return createWeightedRelations(app.list), nil
}
