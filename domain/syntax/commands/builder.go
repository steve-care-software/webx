package commands

import "errors"

type builder struct {
	execution   Execution
	attachment  Attachment
	variable    VariableAssignment
	parameter   ParameterDeclaration
	application ApplicationDeclaration
	module      ModuleDeclaration
}

func createBuilder() Builder {
	out := builder{
		execution:   nil,
		attachment:  nil,
		variable:    nil,
		parameter:   nil,
		application: nil,
		module:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithExecution adds an execution to the builder
func (app *builder) WithExecution(execution Execution) Builder {
	app.execution = execution
	return app
}

// WithAttachment adds an attachment to the builder
func (app *builder) WithAttachment(attachment Attachment) Builder {
	app.attachment = attachment
	return app
}

// WithVariableAssignment adds a variableAssignment to the builder
func (app *builder) WithVariableAssignment(variable VariableAssignment) Builder {
	app.variable = variable
	return app
}

// WithParameterDeclaration adds a parameterDeclaration to the builder
func (app *builder) WithParameterDeclaration(parameter ParameterDeclaration) Builder {
	app.parameter = parameter
	return app
}

// WithApplicationDeclaration adds an applicationDeclaration to the builder
func (app *builder) WithApplicationDeclaration(application ApplicationDeclaration) Builder {
	app.application = application
	return app
}

// WithModuleDeclaration adds a moduleDeclaration to the builder
func (app *builder) WithModuleDeclaration(module ModuleDeclaration) Builder {
	app.module = module
	return app
}

// Now builds a new Command instance
func (app *builder) Now() (Command, error) {
	if app.execution == nil {
		return nil, errors.New("the Execution is mandatory in order to build a Command instance")
	}

	if app.attachment == nil {
		return nil, errors.New("the Attachment is mandatory in order to build a Command instance")
	}

	if app.variable == nil {
		return nil, errors.New("the VariableAssignment is mandatory in order to build a Command instance")
	}

	if app.parameter == nil {
		return nil, errors.New("the ParameterDeclaration is mandatory in order to build a Command instance")
	}

	if app.application == nil {
		return nil, errors.New("the ApplicationDeclaration is mandatory in order to build a Command instance")
	}

	if app.module == nil {
		return nil, errors.New("the ModuleDeclaration is mandatory in order to build a Command instance")
	}

	return createCommand(
		app.execution,
		app.attachment,
		app.variable,
		app.parameter,
		app.application,
		app.module,
	), nil
}
