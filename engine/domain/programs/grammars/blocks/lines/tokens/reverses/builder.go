package reverses

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"

type builder struct {
	escape elements.Element
}

func createBuilder() Builder {
	out := builder{
		escape: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithEscape adds an escape to the builder
func (app *builder) WithEscape(escape elements.Element) Builder {
	app.escape = escape
	return app
}

// Now builds a new Reverse instance
func (app *builder) Now() (Reverse, error) {
	if app.escape != nil {
		return createReverseWithEscape(app.escape), nil
	}

	return createReverse(), nil
}
