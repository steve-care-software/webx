package instructions

import (
	"errors"
)

type referenceBuilder struct {
	element string
	pIndex  *uint
}

func createReferenceBuilder() ReferenceBuilder {
	out := referenceBuilder{
		element: "",
		pIndex:  nil,
	}

	return &out
}

// Create initializes the referenceBuilder
func (app *referenceBuilder) Create() ReferenceBuilder {
	return createReferenceBuilder()
}

// WithElement adds an element to the referenceBuilder
func (app *referenceBuilder) WithElement(element string) ReferenceBuilder {
	app.element = element
	return app
}

// WithIndex adds an index to the referenceBuilder
func (app *referenceBuilder) WithIndex(index uint) ReferenceBuilder {
	app.pIndex = &index
	return app
}

// Now builds a new Reference instance
func (app *referenceBuilder) Now() (Reference, error) {
	if app.element == "" {
		return nil, errors.New("the element is mandatory in order to build a Reference instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Reference instance")
	}

	return createReference(
		app.element,
		*app.pIndex,
	), nil
}
