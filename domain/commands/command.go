package commands

import "github.com/steve-care-software/webx/domain/criterias"

type command struct {
	execution   criterias.Criteria
	attachment  Attachment
	variable    VariableAssignment
	parameter   ParameterDeclaration
	application ApplicationDeclaration
	module      criterias.Criteria
}

func createCommand(
	execution criterias.Criteria,
	attachment Attachment,
	variable VariableAssignment,
	parameter ParameterDeclaration,
	application ApplicationDeclaration,
	module criterias.Criteria,
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

// Execution returns the execution's criteria
func (obj *command) Execution() criterias.Criteria {
	return obj.execution
}

// Attachment returns the attachment
func (obj *command) Attachment() Attachment {
	return obj.attachment
}

// VariableAssignment returns the variable assignment
func (obj *command) VariableAssignment() VariableAssignment {
	return obj.variable
}

// ParameterDeclaration returns the parameter declaration
func (obj *command) ParameterDeclaration() ParameterDeclaration {
	return obj.parameter
}

// ApplicationDeclaration returns the application declaration
func (obj *command) ApplicationDeclaration() ApplicationDeclaration {
	return obj.application
}

// ModuleDeclaration returns the module declaration
func (obj *command) ModuleDeclaration() criterias.Criteria {
	return obj.module
}
