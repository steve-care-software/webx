package compilers

import "errors"

type compositionBuilder struct {
	prefix       []byte
	suffix       []byte
	pattern      []byte
	replacements Replacements
}

func createCompositionBuilder() CompositionBuilder {
	out := compositionBuilder{
		prefix:       nil,
		suffix:       nil,
		pattern:      nil,
		replacements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *compositionBuilder) Create() CompositionBuilder {
	return createCompositionBuilder()
}

// WithPrefix adds a prefix to the builder
func (app *compositionBuilder) WithPrefix(prefix []byte) CompositionBuilder {
	app.prefix = prefix
	return app
}

// WithSuffix adds a suffix to the builder
func (app *compositionBuilder) WithSuffix(suffix []byte) CompositionBuilder {
	app.suffix = suffix
	return app
}

// WithPattern adds a pattern to the builder
func (app *compositionBuilder) WithPattern(pattern []byte) CompositionBuilder {
	app.pattern = pattern
	return app
}

// WithReplacements add replacements to the builder
func (app *compositionBuilder) WithReplacements(replacements Replacements) CompositionBuilder {
	app.replacements = replacements
	return app
}

// Now builds a new Composition instance
func (app *compositionBuilder) Now() (Composition, error) {
	if app.prefix != nil && len(app.prefix) <= 0 {
		app.prefix = nil
	}

	if app.suffix != nil && len(app.suffix) <= 0 {
		app.suffix = nil
	}

	if app.pattern != nil && len(app.pattern) <= 0 {
		app.pattern = nil
	}

	if app.prefix == nil {
		return nil, errors.New("the prefix is mandatory in order to build a Composition instance")
	}

	if app.suffix == nil {
		return nil, errors.New("the suffix is mandatory in order to build a Composition instance")
	}

	if app.pattern == nil {
		return nil, errors.New("the pattern is mandatory in order to build a Composition instance")
	}

	if app.replacements == nil {
		return nil, errors.New("the replacements is mandatory in order to build a Composition instance")
	}

	return createComposition(app.prefix, app.suffix, app.pattern, app.replacements), nil
}
