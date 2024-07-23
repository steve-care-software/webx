package applications

import (
	"github.com/steve-care-software/webx/engine/vms/applications/binaries"
	"github.com/steve-care-software/webx/engine/vms/applications/instructions"
	execution_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	execution_results "github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results"
	execution_success "github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success"
	execution_outputs "github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	instructionsApp instructions.Application,
	binaryApp binaries.Application,
) Application {
	stackFactory := stacks.NewFactory()
	stackBuilder := stacks.NewBuilder()
	framesBuilder := stacks.NewFramesBuilder()
	frameBuilder := stacks.NewFrameBuilder()
	assignmentsBuilder := stacks.NewAssignmentsBuilder()
	assignmentBuilder := stacks.NewAssignmentBuilder()
	assignableBuilder := stacks.NewAssignableBuilder()
	layerExecutionBuilder := execution_layers.NewExecutionBuilder()
	resultBuilder := execution_results.NewBuilder()
	successBuilder := execution_success.NewBuilder()
	outputBuilder := execution_outputs.NewBuilder()
	return createApplication(
		instructionsApp,
		binaryApp,
		stackFactory,
		stackBuilder,
		framesBuilder,
		frameBuilder,
		assignmentsBuilder,
		assignmentBuilder,
		assignableBuilder,
		layerExecutionBuilder,
		resultBuilder,
		successBuilder,
		outputBuilder,
	)
}

// Factory represents the application factory
type Factory interface {
	Create() Application
}

// Application represents a layer application
type Application interface {
	Execute(layer layers.Layer) (execution_layers.Execution, error)
	ExecuteWithInput(layer layers.Layer, input []byte) (execution_layers.Execution, error)
}
