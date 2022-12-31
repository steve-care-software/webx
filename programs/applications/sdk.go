package applications

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

// NameBytesToString converts a name []byte to a string
type NameBytesToString func(name []byte) string

// NewApplication creates a new application
func NewApplication(
	nameBytesToStringFn NameBytesToString,
) Application {
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
		nameBytesToStringFn,
	)
}

// Application represents a program application
type Application interface {
	Compile(modules modules.Modules, instructions instructions.Instructions) (programs.Program, error)
	Execute(input map[uint]interface{}, program programs.Program) (map[uint]interface{}, error)
}
