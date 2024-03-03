package resources

import "errors"

type builderInstructionBuilder struct {
	method        string
	containsParam bool
}

func createBuilderInstructionBuilder() BuilderInstructionBuilder {
	out := builderInstructionBuilder{
		method:        "",
		containsParam: false,
	}

	return &out
}

// Create initializes the builder
func (app *builderInstructionBuilder) Create() BuilderInstructionBuilder {
	return createBuilderInstructionBuilder()
}

// WithMethod adds a method to the builder
func (app *builderInstructionBuilder) WithMethod(method string) BuilderInstructionBuilder {
	app.method = method
	return app
}

// ContainsParam flags the builder as containing a param
func (app *builderInstructionBuilder) ContainsParam() BuilderInstructionBuilder {
	app.containsParam = true
	return app
}

// Now builds a new BuilderInstruction instance
func (app *builderInstructionBuilder) Now() (BuilderInstruction, error) {
	if app.method == "" {
		return nil, errors.New("the method is mandatory in order to build a BuilderInstruction instance")
	}

	return createBuilderInstruction(app.method, app.containsParam), nil
}
