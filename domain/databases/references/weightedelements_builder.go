package references

import "errors"

type weightedElementsBuilder struct {
	list []WeightedElement
}

func createWeightedElementsBuilder() WeightedElementsBuilder {
	out := weightedElementsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *weightedElementsBuilder) Create() WeightedElementsBuilder {
	return createWeightedElementsBuilder()
}

// WithList adds a list to the builder
func (app *weightedElementsBuilder) WithList(list []WeightedElement) WeightedElementsBuilder {
	app.list = list
	return app
}

// Now builds a new WeightedElements instance
func (app *weightedElementsBuilder) Now() (WeightedElements, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 WeightedElement in order to build an WeightedElements instance")
	}

	return createWeightedElements(app.list), nil
}
