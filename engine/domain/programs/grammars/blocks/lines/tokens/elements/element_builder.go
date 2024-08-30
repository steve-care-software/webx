package elements

import "errors"

type elementBuilder struct {
	rule     string
	block    string
	spacer   string
	constant string
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		rule:     "",
		block:    "",
		spacer:   "",
		constant: "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithRule adds a rule to the builder
func (app *elementBuilder) WithRule(rule string) ElementBuilder {
	app.rule = rule
	return app
}

// WithBlock adds a block to the builder
func (app *elementBuilder) WithBlock(block string) ElementBuilder {
	app.block = block
	return app
}

// WithSpacer adds a spacer to the builder
func (app *elementBuilder) WithSpacer(spacer string) ElementBuilder {
	app.spacer = spacer
	return app
}

// WithConstant adds a constant to the builder
func (app *elementBuilder) WithConstant(constant string) ElementBuilder {
	app.constant = constant
	return app
}

// Now builds a new Element
func (app *elementBuilder) Now() (Element, error) {
	if app.rule != "" {
		return createElementWithRule(app.rule), nil
	}

	if app.block != "" {
		return createElementWithBlock(app.block), nil
	}

	if app.spacer != "" {
		return createElementWithSpacer(app.spacer), nil
	}

	if app.constant != "" {
		return createElementWithConstant(app.constant), nil
	}

	return nil, errors.New("the Element is invalid")
}
