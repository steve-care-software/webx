package entities

import "errors"

type identifierBuilder struct {
	pSection *uint
	pElement *uint
}

func createIdentifierBuilder() IdentifierBuilder {
	out := identifierBuilder{
		pSection: nil,
		pElement: nil,
	}

	return &out
}

// Create initializes the builder
func (app *identifierBuilder) Create() IdentifierBuilder {
	return createIdentifierBuilder()
}

// WithSection adds a section to the builder
func (app *identifierBuilder) WithSection(section uint) IdentifierBuilder {
	app.pSection = &section
	return app
}

// WithElement adds an element to the builder
func (app *identifierBuilder) WithElement(element uint) IdentifierBuilder {
	app.pElement = &element
	return app
}

// Now builds a new Identifier instance
func (app *identifierBuilder) Now() (Identifier, error) {
	if app.pSection == nil {
		return nil, errors.New("the section is mandatory in order to build an Identifier instance")
	}

	if app.pElement == nil {
		return nil, errors.New("the element is mandatory in order to build an Identifier instance")
	}

	return createIdentifier(*app.pSection, *app.pElement), nil
}
