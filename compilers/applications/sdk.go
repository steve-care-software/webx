package applications

import (
	"github.com/steve-care-software/webx/compilers/domain/compilers"
	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	selector_applications "github.com/steve-care-software/webx/selectors/applications"
)

// NewApplication creates a new application instance
func NewApplication() Application {
	grammarApp := grammar_applications.NewApplication()
	selectorApp := selector_applications.NewApplication()
	programBuilder := programs.NewBuilder()
	instructionsBuilder := programs.NewInstructionsBuilder()
	instructionBuilder := programs.NewInstructionBuilder()
	applicationBuilder := programs.NewApplicationBuilder()
	attachmentsBuilder := programs.NewAttachmentsBuilder()
	attachmentBuilder := programs.NewAttachmentBuilder()
	valueBuilder := programs.NewValueBuilder()
	return createApplication(
		grammarApp,
		selectorApp,
		programBuilder,
		instructionsBuilder,
		instructionBuilder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		valueBuilder,
	)
}

// Application represents the compiler application
type Application interface {
	Execute(compiler compilers.Compiler, modules modules.Modules, script []byte) (programs.Program, error)
}
