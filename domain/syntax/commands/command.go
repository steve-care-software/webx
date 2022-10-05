package commands

type command struct {
	execution   Execution
	attachment  Attachment
	variable    VariableAssignment
	parameter   ParameterDeclaration
	application ApplicationDeclaration
	module      ModuleDeclaration
}

func createCommand(
	execution Execution,
	attachment Attachment,
	variable VariableAssignment,
	parameter ParameterDeclaration,
	application ApplicationDeclaration,
	module ModuleDeclaration,
) Command {
	out := command{
		execution:   execution,
		attachment:  attachment,
		variable:    variable,
		parameter:   parameter,
		application: application,
		module:      module,
	}

	return &out
}

// Execution returns the execution, if any
func (obj *command) Execution() Execution {
	return obj.execution
}

// Attachment returns the attachment, if any
func (obj *command) Attachment() Attachment {
	return obj.attachment
}

// VariableAssignment returns the variable assignment, if any
func (obj *command) VariableAssignment() VariableAssignment {
	return obj.variable
}

// ParameterDeclaration returns the parameter declaration, if any
func (obj *command) ParameterDeclaration() ParameterDeclaration {
	return obj.parameter
}

// ApplicationDeclaration returns the application declaration, if any
func (obj *command) ApplicationDeclaration() ApplicationDeclaration {
	return obj.application
}

// ModuleDeclaration returns the module declaration, if any
func (obj *command) ModuleDeclaration() ModuleDeclaration {
	return obj.module
}
