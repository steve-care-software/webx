package programs

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/programs"
	"github.com/steve-care-software/webx/domain/programs/modules"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	builder := programs.NewBuilder()
	instructionBuilder := programs.NewInstructionBuilder()
	applicationBuilder := programs.NewApplicationBuilder()
	attachmentsBuilder := programs.NewAttachmentsBuilder()
	attachmentBuilder := programs.NewAttachmentBuilder()
	assignmentBuilder := programs.NewAssignmentBuilder()
	valueBuilder := programs.NewValueBuilder()
	return createBuilder(
		hashAdapter,
		builder,
		instructionBuilder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		assignmentBuilder,
		valueBuilder,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithModules(modules modules.Modules) Builder
	Now() (Application, error)
}

// Application represents a program application
type Application interface {
	Execute(instructions instructions.Instructions) (programs.Program, error)
}
