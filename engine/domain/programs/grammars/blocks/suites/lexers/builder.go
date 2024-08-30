package lexers

import "errors"

type builder struct {
	output []byte
	isFail bool
}

func createBuilder() Builder {
	out := builder{
		output: nil,
		isFail: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithOutput adds an output to the builder
func (app *builder) WithOutput(output []byte) Builder {
	app.output = output
	return app
}

// IsFail flags the builder as fail
func (app *builder) IsFail() Builder {
	app.isFail = true
	return app
}

// Now builds a new Lexer instance
func (app *builder) Now() (Lexer, error) {
	if app.output != nil && len(app.output) <= 0 {
		app.output = nil
	}

	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a Lexer instance")
	}

	return createLexer(
		app.output,
		app.isFail,
	), nil
}
