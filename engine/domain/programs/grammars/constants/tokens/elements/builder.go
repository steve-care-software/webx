package elements

import "errors"

type builder struct {
	rule     string
	constant string
}

func createBuilder() Builder {
	out := builder{
		rule:     "",
		constant: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRule adds a rule to the builder
func (app *builder) WithRule(rule string) Builder {
	app.rule = rule
	return app
}

// WithConstant adds a constant to the builder
func (app *builder) WithConstant(constant string) Builder {
	app.constant = constant
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
	if app.rule != "" {
		return createElementWithRule(app.rule), nil
	}

	if app.constant != "" {
		return createElementWithConstant(app.constant), nil
	}

	return nil, errors.New("the Element is invalid")
}
