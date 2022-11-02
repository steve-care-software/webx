package programs

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/programs"
	"github.com/steve-care-software/webx/domain/programs/modules"
)

type builder struct {
	hashAdapter        hash.Adapter
	builder            programs.Builder
	instructionBuilder programs.InstructionBuilder
	applicationBuilder programs.ApplicationBuilder
	attachmentsBuilder programs.AttachmentsBuilder
	attachmentBuilder  programs.AttachmentBuilder
	assignmentBuilder  programs.AssignmentBuilder
	valueBuilder       programs.ValueBuilder
	modules            modules.Modules
}

func createBuilder(
	hashAdapter hash.Adapter,
	builderIns programs.Builder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	assignmentBuilder programs.AssignmentBuilder,
	valueBuilder programs.ValueBuilder,
) Builder {
	out := builder{
		hashAdapter:        hashAdapter,
		builder:            builderIns,
		instructionBuilder: instructionBuilder,
		applicationBuilder: applicationBuilder,
		attachmentsBuilder: attachmentsBuilder,
		attachmentBuilder:  attachmentBuilder,
		assignmentBuilder:  assignmentBuilder,
		valueBuilder:       valueBuilder,
		modules:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
		app.builder,
		app.instructionBuilder,
		app.applicationBuilder,
		app.attachmentsBuilder,
		app.attachmentBuilder,
		app.assignmentBuilder,
		app.valueBuilder,
	)
}

// WithModules add modules to the builder
func (app *builder) WithModules(modules modules.Modules) Builder {
	app.modules = modules
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.modules != nil {
		return createApplicationWithModules(
			app.hashAdapter,
			app.builder,
			app.instructionBuilder,
			app.applicationBuilder,
			app.attachmentsBuilder,
			app.attachmentBuilder,
			app.assignmentBuilder,
			app.valueBuilder,
			app.modules,
		), nil
	}

	return createApplication(
		app.hashAdapter,
		app.builder,
		app.instructionBuilder,
		app.applicationBuilder,
		app.attachmentsBuilder,
		app.attachmentBuilder,
		app.assignmentBuilder,
		app.valueBuilder,
	), nil
}
