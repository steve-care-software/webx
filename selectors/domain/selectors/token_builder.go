package selectors

import "errors"

type tokenBuilder struct {
	name        string
	reverseName string
	element     Element
	pContent    *uint
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:        "",
		reverseName: "",
		element:     nil,
		pContent:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithName adds a name to the builder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// WithReverseName adds a reverseName to the builder
func (app *tokenBuilder) WithReverseName(reverseName string) TokenBuilder {
	app.reverseName = reverseName
	return app
}

// WithElement adds an element to the builder
func (app *tokenBuilder) WithElement(element Element) TokenBuilder {
	app.element = element
	return app
}

// WithContent adds a content index to the builder
func (app *tokenBuilder) WithContent(content uint) TokenBuilder {
	app.pContent = &content
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.reverseName == "" {
		return nil, errors.New("the reverseName is mandatory in order to build a Token instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	if app.pContent != nil {
		return createTokenWithContentIndex(app.name, app.reverseName, app.element, app.pContent), nil
	}

	return createToken(app.name, app.reverseName, app.element), nil
}
