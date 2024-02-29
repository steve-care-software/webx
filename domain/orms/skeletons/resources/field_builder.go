package resources

import "errors"

type fieldBuilder struct {
	name      string
	retriever []string
	builder   string
	condition string
	kind      Kind
	canBeNil  bool
}

func createFieldBuilder() FieldBuilder {
	out := fieldBuilder{
		name:      "",
		retriever: nil,
		builder:   "",
		condition: "",
		kind:      nil,
		canBeNil:  false,
	}

	return &out
}

// Create initializes the builder
func (app *fieldBuilder) Create() FieldBuilder {
	return createFieldBuilder()
}

// WithName adds a name to the builder
func (app *fieldBuilder) WithName(name string) FieldBuilder {
	app.name = name
	return app
}

// WithRetriever adds a retriever to the builder
func (app *fieldBuilder) WithRetriever(retriever []string) FieldBuilder {
	app.retriever = retriever
	return app
}

// WithBuilder adds a builder to the builder
func (app *fieldBuilder) WithBuilder(builder string) FieldBuilder {
	app.builder = builder
	return app
}

// WithCondition adds a condition to the builder
func (app *fieldBuilder) WithCondition(condition string) FieldBuilder {
	app.condition = condition
	return app
}

// WithKind adds a kind to the builder
func (app *fieldBuilder) WithKind(kind Kind) FieldBuilder {
	app.kind = kind
	return app
}

// CanBeNil flags the builder as canBeNil
func (app *fieldBuilder) CanBeNil() FieldBuilder {
	app.canBeNil = true
	return app
}

// Now builds a new Field instance
func (app *fieldBuilder) Now() (Field, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Field instance")
	}

	if app.retriever != nil && len(app.retriever) <= 0 {
		app.retriever = nil
	}

	if app.retriever == nil {
		return nil, errors.New("the retriever method is mandatory in order to build a Field instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Field instance")
	}

	if app.builder != "" && app.condition != "" {
		return createFieldWithConditionAndBuilder(
			app.name,
			app.retriever,
			app.kind,
			app.canBeNil,
			app.condition,
			app.builder,
		), nil
	}

	if app.builder != "" {
		return createFieldWithBuilder(
			app.name,
			app.retriever,
			app.kind,
			app.canBeNil,
			app.builder,
		), nil
	}

	if app.condition != "" {
		return createFieldWithCondition(
			app.name,
			app.retriever,
			app.kind,
			app.canBeNil,
			app.condition,
		), nil
	}

	return createField(
		app.name,
		app.retriever,
		app.kind,
		app.canBeNil,
	), nil
}
