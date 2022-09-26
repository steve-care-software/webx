package commands

type content struct {
	execution   Execution
	attachment  Attachment
	variable    VariableAssignment
	parameter   ParameterDeclaration
	application ApplicationDeclaration
	module      ModuleDeclaration
}

func createContentWithExecution(
	execution Execution,
) Content {
	return createContentInternally(execution, nil, nil, nil, nil, nil)
}

func createContentWithAttachment(
	attachment Attachment,
) Content {
	return createContentInternally(nil, attachment, nil, nil, nil, nil)
}

func createContentWithVariableAssignment(
	variable VariableAssignment,
) Content {
	return createContentInternally(nil, nil, variable, nil, nil, nil)
}

func createContentWithParameterDeclaration(
	parameter ParameterDeclaration,
) Content {
	return createContentInternally(nil, nil, nil, parameter, nil, nil)
}

func createContentWithApplicationDeclaration(
	application ApplicationDeclaration,
) Content {
	return createContentInternally(nil, nil, nil, nil, application, nil)
}

func createContentWithModuleDeclaration(
	module ModuleDeclaration,
) Content {
	return createContentInternally(nil, nil, nil, nil, nil, module)
}

func createContentInternally(
	execution Execution,
	attachment Attachment,
	variable VariableAssignment,
	parameter ParameterDeclaration,
	application ApplicationDeclaration,
	module ModuleDeclaration,
) Content {
	out := content{
		execution:   execution,
		attachment:  attachment,
		variable:    variable,
		parameter:   parameter,
		application: application,
		module:      module,
	}

	return &out
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *content) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() Execution {
	return obj.execution
}

// IsAttachment returns true if there is an attachment, false otherwise
func (obj *content) IsAttachment() bool {
	return obj.attachment != nil
}

// Attachment returns the attachment, if any
func (obj *content) Attachment() Attachment {
	return obj.attachment
}

// IsVariableAssignment returns true if there is an variable assignment, false otherwise
func (obj *content) IsVariableAssignment() bool {
	return obj.variable != nil
}

// VariableAssignment returns the variable assignment, if any
func (obj *content) VariableAssignment() VariableAssignment {
	return obj.variable
}

// IsParameterDeclaration returns true if there is an parameter declaration, false otherwise
func (obj *content) IsParameterDeclaration() bool {
	return obj.parameter != nil
}

// ParameterDeclaration returns the parameter declaration, if any
func (obj *content) ParameterDeclaration() ParameterDeclaration {
	return obj.parameter
}

// IsApplicationDeclaration returns true if there is an application declaration, false otherwise
func (obj *content) IsApplicationDeclaration() bool {
	return obj.application != nil
}

// ApplicationDeclaration returns the application declaration, if any
func (obj *content) ApplicationDeclaration() ApplicationDeclaration {
	return obj.application
}

// IsModuleDeclaration returns true if there is a module declaration, false otherwise
func (obj *content) IsModuleDeclaration() bool {
	return obj.module != nil
}

// ModuleDeclaration returns the module declaration, if any
func (obj *content) ModuleDeclaration() ModuleDeclaration {
	return obj.module
}
