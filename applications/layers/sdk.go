package layers

import (
	applications_instructions "github.com/steve-care-software/datastencil/applications/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/commands"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/success"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execInsApp applications_instructions.Application,
	tempBaseDir string,
) Application {
	stackFactory := stacks.NewFactory()
	stackBuilder := stacks.NewBuilder()
	framesBuilder := stacks.NewFramesBuilder()
	frameBuilder := stacks.NewFrameBuilder()
	assignmentsBuilder := stacks.NewAssignmentsBuilder()
	assignmentBuilder := stacks.NewAssignmentBuilder()
	assignableBuilder := stacks.NewAssignableBuilder()
	resultBuilder := results.NewBuilder()
	successBuilder := success.NewBuilder()
	outputBuilder := outputs.NewBuilder()
	return createApplication(
		execInsApp,
		stackFactory,
		stackBuilder,
		framesBuilder,
		frameBuilder,
		assignmentsBuilder,
		assignmentBuilder,
		assignableBuilder,
		resultBuilder,
		successBuilder,
		outputBuilder,
		tempBaseDir,
	)
}

// Application represents an application
type Application interface {
	Execute(input []byte, layer layers.Layer, context commands.Commands) (results.Result, error)
}
