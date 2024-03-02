package resources

import "errors"

type kindBuilder struct {
	native     Native
	reference  []string
	connection string
}

func createKindBuilder() KindBuilder {
	out := kindBuilder{
		native:     nil,
		reference:  nil,
		connection: "",
	}

	return &out
}

// Create initializes the builder
func (app *kindBuilder) Create() KindBuilder {
	return createKindBuilder()
}

// WithNative adds a native to the builder
func (app *kindBuilder) WithNative(native Native) KindBuilder {
	app.native = native
	return app
}

// WithReference adds a reference to the builder
func (app *kindBuilder) WithReference(reference []string) KindBuilder {
	app.reference = reference
	return app
}

// WithConnection adds a connection to the builder
func (app *kindBuilder) WithConnection(connection string) KindBuilder {
	app.connection = connection
	return app
}

// Now builds a new Kind isntance
func (app *kindBuilder) Now() (Kind, error) {
	if app.reference != nil && len(app.reference) <= 0 {
		app.reference = nil
	}

	if app.native != nil {
		return createKindWithNative(app.native), nil
	}

	if app.reference != nil {
		return createKindWithReference(app.reference), nil
	}

	if app.connection != "" {
		return createKindWithConnection(app.connection), nil
	}

	return nil, errors.New("the Kind is invalid")
}
