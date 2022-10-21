package instructions

import "errors"

type outputBuilder struct {
	instructions Instructions
	remaining    []byte
}

func createOutputBuilder() OutputBuilder {
	out := outputBuilder{
		instructions: nil,
		remaining:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *outputBuilder) Create() OutputBuilder {
	return createOutputBuilder()
}

// WithInstructions add instructions to the builder
func (app *outputBuilder) WithInstructions(instructions Instructions) OutputBuilder {
	app.instructions = instructions
	return app
}

// WithRemaining add remaining to the builder
func (app *outputBuilder) WithRemaining(remaining []byte) OutputBuilder {
	app.remaining = remaining
	return app
}

// Now builds a new Output instance
func (app *outputBuilder) Now() (Output, error) {
	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build an Output instance")
	}

	if app.remaining != nil && len(app.remaining) <= 0 {
		app.remaining = nil
	}

	if app.remaining != nil {
		return createOutputWithRemaining(app.instructions, app.remaining), nil
	}

	return createOutput(app.instructions), nil
}
