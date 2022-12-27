package applications

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/programs"
)

// NewApplication creates a new application
func NewApplication() Application {
	builder := programs.NewBuilder()
	instructionsBuilder := programs.NewInstructionsBuilder()
	instructionBuilder := programs.NewInstructionBuilder()
	applicationBuilder := programs.NewApplicationBuilder()
	attachmentsBuilder := programs.NewAttachmentsBuilder()
	attachmentBuilder := programs.NewAttachmentBuilder()
	valueBuilder := programs.NewValueBuilder()
	hashAdapter := hash.NewAdapter()
	return createApplication(
		builder,
		instructionsBuilder,
		instructionBuilder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		valueBuilder,
		hashAdapter,
	)
}

// Application represents a program application
type Application interface {
	Execute(input map[uint]interface{}, program programs.Program) (map[uint]interface{}, error)
}
