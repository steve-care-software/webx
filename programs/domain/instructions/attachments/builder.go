package attachments

import "errors"

type builder struct {
	variable    Variable
	application []byte
}

func createBuilder() Builder {
	out := builder{
		variable:    nil,
		application: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable Variable) Builder {
	app.variable = variable
	return app
}

// WithApplication adds an application to the builder
func (app *builder) WithApplication(application []byte) Builder {
	app.application = application
	return app
}

// Now builds a new Attachment instance
func (app *builder) Now() (Attachment, error) {
	if app.variable == nil {
		return nil, errors.New("the variable is mandatory in order to build an Attachment instance")
	}

	if app.application == nil {
		return nil, errors.New("the application is mandatory in order to build an Attachment instance")
	}

	return createAttachment(app.variable, app.application), nil
}
