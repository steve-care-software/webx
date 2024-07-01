package layers

import (
	"github.com/steve-care-software/datastencil/applications/links/layers/binaries"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions"
	execution_layers "github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	instructionsApp    instructions.Application
	binaryApp          binaries.Application
	stackFactory       stacks.Factory
	stackBuilder       stacks.Builder
	framesBuilder      stacks.FramesBuilder
	frameBuilder       stacks.FrameBuilder
	assignmentsBuilder stacks.AssignmentsBuilder
	assignmentBuilder  stacks.AssignmentBuilder
	assignableBuilder  stacks.AssignableBuilder
	layerBuilder       execution_layers.LayerBuilder
}

func createApplication(
	instructionsApp instructions.Application,
	binaryApp binaries.Application,
	stackFactory stacks.Factory,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	assignmentBuilder stacks.AssignmentBuilder,
	assignableBuilder stacks.AssignableBuilder,
	layerBuilder execution_layers.LayerBuilder,
	resultBuilder results.Builder,
) Application {
	out := application{
		instructionsApp:    instructionsApp,
		binaryApp:          binaryApp,
		stackFactory:       stackFactory,
		stackBuilder:       stackBuilder,
		framesBuilder:      framesBuilder,
		frameBuilder:       frameBuilder,
		assignmentsBuilder: assignmentsBuilder,
		assignmentBuilder:  assignmentBuilder,
		assignableBuilder:  assignableBuilder,
		layerBuilder:       layerBuilder,
	}

	return &out
}

// ExecuteWithInput executes the layer using the provided input and returns the executions
func (app *application) ExecuteWithInput(layer layers.Layer, input []byte) (execution_layers.Layer, error) {
	return app.execute(layer, input)
}

// Execute executes a layer
func (app *application) Execute(layer layers.Layer) (execution_layers.Layer, error) {
	return app.execute(layer, nil)
}

func (app *application) execute(layer layers.Layer, input []byte) (execution_layers.Layer, error) {
	stack, err := app.stack(layer, input)
	if err != nil {
		return nil, err
	}

	instructions := layer.Instructions()
	result, err := app.instructionsApp.Execute(stack, instructions)
	if err != nil {
		return nil, err
	}

	return app.layerBuilder.Create().
		WithInput(input).
		WithSource(layer).
		WithResult(result).
		Now()
}

func (app *application) stack(layer layers.Layer, input []byte) (stacks.Stack, error) {
	if input != nil && len(input) <= 0 {
		input = nil
	}

	if input == nil {
		return app.stackFactory.Create(), nil
	}

	assignable, err := app.assignableBuilder.Create().WithBytes(input).Now()
	if err != nil {
		return nil, err
	}

	inputVariable := layer.Input()
	assignment, err := app.assignmentBuilder.Create().WithAssignable(assignable).WithName(inputVariable).Now()
	if err != nil {
		return nil, err
	}

	assignments, err := app.assignmentsBuilder.Create().WithList([]stacks.Assignment{
		assignment,
	}).Now()

	if err != nil {
		return nil, err
	}

	frame, err := app.frameBuilder.Create().WithAssignments(assignments).Now()
	if err != nil {
		return nil, err
	}

	frames, err := app.framesBuilder.Create().WithList([]stacks.Frame{
		frame,
	}).Now()

	if err != nil {
		return nil, err
	}

	return app.stackBuilder.Create().WithFrames(frames).Now()
}
