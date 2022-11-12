package references

import "errors"

type weightedElementBuilder struct {
	pIndex  *uint
	pWeight *uint
}

func createWeightedElementBuilder() WeightedElementBuilder {
	out := weightedElementBuilder{
		pIndex:  nil,
		pWeight: nil,
	}

	return &out
}

// Create initializes the builder
func (app *weightedElementBuilder) Create() WeightedElementBuilder {
	return createWeightedElementBuilder()
}

// WithIndex adds an index to the builder
func (app *weightedElementBuilder) WithIndex(index uint) WeightedElementBuilder {
	app.pIndex = &index
	return app
}

// WithWeight adds a weight to the builder
func (app *weightedElementBuilder) WithWeight(weight uint) WeightedElementBuilder {
	app.pWeight = &weight
	return app
}

// Now builds a new WeightedElement instance
func (app *weightedElementBuilder) Now() (WeightedElement, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a WeightedElement instance")
	}

	if app.pWeight == nil {
		return nil, errors.New("the weight is mandatory in order to build a WeightedElement instance")
	}

	return createWeightedElement(*app.pIndex, *app.pWeight), nil
}
