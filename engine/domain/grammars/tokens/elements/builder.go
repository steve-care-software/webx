package elements

import "errors"

type builder struct {
	rule  string
	block string
}

func createBuilder() Builder {
	out := builder{
		rule:  "",
		block: "",
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

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block string) Builder {
	app.block = block
	return app
}

// Now builds a new Element
func (app *builder) Now() (Element, error) {
	if app.rule != "" {
		return createElementWithRule(app.rule), nil
	}

	if app.block != "" {
		return createElementWithBlock(app.block), nil
	}

	return nil, errors.New("the Element is invalid")
}
