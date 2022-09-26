package commands

import "errors"

type contentBuilder struct {
	execution   Execution
	attachment  Attachment
	variable    VariableAssignment
	parameter   ParameterDeclaration
	application ApplicationDeclaration
	module      ModuleDeclaration
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
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
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// WithExecution adds an execution to the builder
func (app *contentBuilder) WithExecution(execution Execution) ContentBuilder {
	app.execution = execution
	return app
}

// WithAttachment adds an attachment to the builder
func (app *contentBuilder) WithAttachment(attachment Attachment) ContentBuilder {
	app.attachment = attachment
	return app
}

// WithVariableAssignment adds a variableAssignment to the builder
func (app *contentBuilder) WithVariableAssignment(variable VariableAssignment) ContentBuilder {
	app.variable = variable
	return app
}

// WithParameterDeclaration adds a parameterDeclaration to the builder
func (app *contentBuilder) WithParameterDeclaration(parameter ParameterDeclaration) ContentBuilder {
	app.parameter = parameter
	return app
}

// WithApplicationDeclaration adds an applicationDeclaration to the builder
func (app *contentBuilder) WithApplicationDeclaration(application ApplicationDeclaration) ContentBuilder {
	app.application = application
	return app
}

// WithModuleDeclaration adds a moduleDeclaration to the builder
func (app *contentBuilder) WithModuleDeclaration(module ModuleDeclaration) ContentBuilder {
	app.module = module
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.execution != nil {
		return createContentWithExecution(app.execution), nil
	}

	if app.attachment != nil {
		return createContentWithAttachment(app.attachment), nil
	}

	if app.variable != nil {
		return createContentWithVariableAssignment(app.variable), nil
	}

	if app.parameter != nil {
		return createContentWithParameterDeclaration(app.parameter), nil
	}

	if app.application != nil {
		return createContentWithApplicationDeclaration(app.application), nil
	}

	if app.module != nil {
		return createContentWithModuleDeclaration(app.module), nil
	}

	return nil, errors.New("the Content is invalid")
}
