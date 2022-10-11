package commands

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	return createExecutionBuilder()
}

// NewAttachmentBuilder creates a new attachment builder
func NewAttachmentBuilder() AttachmentBuilder {
	return createAttachmentBuilder()
}

// NewVariableAssignmentBuilder creates a new variable assignment builder
func NewVariableAssignmentBuilder() VariableAssignmentBuilder {
	return createVariableAssignmentBuilder()
}

// NewParameterDeclarationBuilder creates a new parameter declaration builder
func NewParameterDeclarationBuilder() ParameterDeclarationBuilder {
	return createParameterDeclarationBuilder()
}

// NewApplicationDeclarationBuilder creates a new application declaration builder
func NewApplicationDeclarationBuilder() ApplicationDeclarationBuilder {
	return createApplicationDeclarationBuilder()
}

// NewModuleDeclarationBuilder creates a new module declaration builder
func NewModuleDeclarationBuilder() ModuleDeclarationBuilder {
	return createModuleDeclarationBuilder()
}

// Builder represents a command builder
type Builder interface {
	Create() Builder
	WithExecution(execution Execution) Builder
	WithAttachment(attachment Attachment) Builder
	WithVariableAssignment(variableAssignment VariableAssignment) Builder
	WithParameterDeclaration(parameterDeclaration ParameterDeclaration) Builder
	WithApplicationDeclaration(applicationDeclaration ApplicationDeclaration) Builder
	WithModuleDeclaration(moduleDeclaration ModuleDeclaration) Builder
	Now() (Command, error)
}

// Command represents the possible commands
type Command interface {
	Execution() Execution
	Attachment() Attachment
	VariableAssignment() VariableAssignment
	ParameterDeclaration() ParameterDeclaration
	ApplicationDeclaration() ApplicationDeclaration
	ModuleDeclaration() ModuleDeclaration
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithApplication(application criterias.Criteria) ExecutionBuilder
	WithAssignee(assignee criterias.Criteria) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Application() criterias.Criteria
	Assignee() criterias.Criteria
}

// AttachmentBuilder represents an attachment builder
type AttachmentBuilder interface {
	Create() AttachmentBuilder
	WithGlobal(global criterias.Criteria) AttachmentBuilder
	WithLocal(local criterias.Criteria) AttachmentBuilder
	WithApplication(application criterias.Criteria) AttachmentBuilder
	Now() (Attachment, error)
}

// Attachment represents an attachment
type Attachment interface {
	Global() criterias.Criteria
	Local() criterias.Criteria
	Application() criterias.Criteria
}

// VariableAssignmentBuilder represents a variable assignment builder
type VariableAssignmentBuilder interface {
	Create() VariableAssignmentBuilder
	WithAssignee(assignee criterias.Criteria) VariableAssignmentBuilder
	WithValue(value criterias.Criteria) VariableAssignmentBuilder
	Now() (VariableAssignment, error)
}

// VariableAssignment represents a variable assignment
type VariableAssignment interface {
	Assignee() criterias.Criteria
	Value() criterias.Criteria
}

// ParameterDeclarationBuilder represents a parameter declaration builder
type ParameterDeclarationBuilder interface {
	Create() ParameterDeclarationBuilder
	WithInput(input criterias.Criteria) ParameterDeclarationBuilder
	WithOutput(output criterias.Criteria) ParameterDeclarationBuilder
	Now() (ParameterDeclaration, error)
}

// ParameterDeclaration represents a parameter declaration
type ParameterDeclaration interface {
	Input() criterias.Criteria
	Output() criterias.Criteria
}

// ApplicationDeclarationBuilder represents an application declaration builder
type ApplicationDeclarationBuilder interface {
	Create() ApplicationDeclarationBuilder
	WithModule(module criterias.Criteria) ApplicationDeclarationBuilder
	WithName(name criterias.Criteria) ApplicationDeclarationBuilder
	Now() (ApplicationDeclaration, error)
}

// ApplicationDeclaration represents an application declaration
type ApplicationDeclaration interface {
	Module() criterias.Criteria
	Name() criterias.Criteria
}

// ModuleDeclarationBuilder represents a module declaration builder
type ModuleDeclarationBuilder interface {
	Create() ModuleDeclarationBuilder
	WithName(name criterias.Criteria) ModuleDeclarationBuilder
	Now() (ModuleDeclaration, error)
}

// ModuleDeclaration represents a module declaration
type ModuleDeclaration interface {
	Name() criterias.Criteria
}
