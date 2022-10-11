package programs

import (
	"github.com/steve-care-software/syntax/applications/engines/criterias"
	grammar_application "github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	grammarApp := grammar_application.NewApplication()
	criteriaApp := criterias.NewApplication()
	builder := programs.NewBuilder()
	applicationBuilder := applications.NewBuilder()
	attachmentsBuilder := applications.NewAttachmentsBuilder()
	attachmentBuilder := applications.NewAttachmentBuilder()
	assignmentBuilder := applications.NewAssignmentBuilder()
	valueBuilder := applications.NewValueBuilder()
	return createBuilder(grammarApp, criteriaApp, builder, applicationBuilder, attachmentsBuilder, attachmentBuilder, assignmentBuilder, valueBuilder)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithModules(modules modules.Modules) Builder
	Now() (Application, error)
}

// Application represents a program application
type Application interface {
	Execute(grammar grammars.Grammar, command commands.Command, script []byte) (programs.Program, []byte, error)
}
