package replacements

import "errors"

type replacementBuilder struct {
	origin string
	target string
}

func createReplacementBuilder() ReplacementBuilder {
	out := replacementBuilder{
		origin: "",
		target: "",
	}

	return &out
}

// Create initializes the builder
func (app *replacementBuilder) Create() ReplacementBuilder {
	return createReplacementBuilder()
}

// WithOrigin adds an origin to the builder
func (app *replacementBuilder) WithOrigin(origin string) ReplacementBuilder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *replacementBuilder) WithTarget(target string) ReplacementBuilder {
	app.target = target
	return app
}

// Now builds a new Replacement instance
func (app *replacementBuilder) Now() (Replacement, error) {
	if app.origin == "" {
		return nil, errors.New("the origin is mandatory in order to build a Replacement instance")
	}

	if app.target == "" {
		return nil, errors.New("the target is mandatory in order to build a Replacement instance")
	}

	return createReplacement(
		app.origin,
		app.target,
	), nil
}
