package trees

import "errors"

type elementsBuilder struct {
	list []Element
}

func createElementsBuilder() ElementsBuilder {
	out := elementsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementsBuilder) Create() ElementsBuilder {
	return createElementsBuilder()
}

// WithList adds a list to the builder
func (app *elementsBuilder) WithList(list []Element) ElementsBuilder {
	app.list = list
	return app
}

// Now builds a new Elements instance
func (app *elementsBuilder) Now() (Elements, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Element in order to build a Elements instance")
	}

	return createElements(app.list), nil
}
