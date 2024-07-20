package layers

import (
	"github.com/steve-care-software/datastencil/applications/layers/binaries"
	"github.com/steve-care-software/datastencil/applications/layers/instructions"
	execution_layers "github.com/steve-care-software/datastencil/domain/instances/executions"
	execution_results "github.com/steve-care-software/datastencil/domain/instances/executions/results"
	"github.com/steve-care-software/datastencil/domain/instances/executions/results/success"
	execution_success "github.com/steve-care-software/datastencil/domain/instances/executions/results/success"
	execution_outputs "github.com/steve-care-software/datastencil/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	instructionsApp       instructions.Application
	binaryApp             binaries.Application
	stackFactory          stacks.Factory
	stackBuilder          stacks.Builder
	framesBuilder         stacks.FramesBuilder
	frameBuilder          stacks.FrameBuilder
	assignmentsBuilder    stacks.AssignmentsBuilder
	assignmentBuilder     stacks.AssignmentBuilder
	assignableBuilder     stacks.AssignableBuilder
	layerExecutionBuilder execution_layers.ExecutionBuilder
	resultBuilder         execution_results.Builder
	successBuilder        execution_success.Builder
	outputBuilder         execution_outputs.Builder
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
	layerExecutionBuilder execution_layers.ExecutionBuilder,
	resultBuilder execution_results.Builder,
) Application {
	out := application{
		instructionsApp:       instructionsApp,
		binaryApp:             binaryApp,
		stackFactory:          stackFactory,
		stackBuilder:          stackBuilder,
		framesBuilder:         framesBuilder,
		frameBuilder:          frameBuilder,
		assignmentsBuilder:    assignmentsBuilder,
		assignmentBuilder:     assignmentBuilder,
		assignableBuilder:     assignableBuilder,
		layerExecutionBuilder: layerExecutionBuilder,
		resultBuilder:         resultBuilder,
	}

	return &out
}

// Execute executes a layer
func (app *application) Execute(layer layers.Layer) (execution_layers.Execution, error) {
	return app.execute(layer, nil)
}

// ExecuteWithInput executes the layer using the provided input and returns the executions
func (app *application) ExecuteWithInput(layer layers.Layer, input []byte) (execution_layers.Execution, error) {
	return app.execute(layer, input)
}

func (app *application) execute(layer layers.Layer, input []byte) (execution_layers.Execution, error) {
	frame, err := app.frame(layer, input)
	if err != nil {
		return nil, err
	}

	instructions := layer.Instructions()
	retFrame, retInterruption, err := app.instructionsApp.Execute(frame, instructions)
	if err != nil {
		return nil, err
	}

	resultBuilder := app.resultBuilder.Create()
	if retInterruption != nil {
		resultBuilder.WithInterruption(retInterruption)
	}

	if retFrame != nil {
		output := layer.Output()
		success, err := app.success(retFrame, output)
		if err != nil {
			return nil, err
		}

		resultBuilder.WithSuccess(success)
	}

	result, err := resultBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.layerExecutionBuilder.Create().
		WithInput(input).
		WithSource(layer).
		WithResult(result).
		Now()
}

func (app *application) success(frame stacks.Frame, output outputs.Output) (success.Success, error) {
	variable := output.Variable()
	value, err := frame.FetchBytes(variable)
	if err != nil {
		return nil, err
	}

	outputBuilder := app.outputBuilder.Create().WithInput(value)
	if output.HasExecute() {
		commands := output.Execute()
		retResult, err := app.binaryApp.Execute(value, commands)
		if err != nil {
			return nil, err
		}

		outputBuilder.WithExecute(retResult)
	}

	retOutput, err := outputBuilder.Now()
	if err != nil {
		return nil, err
	}

	kind := output.Kind()
	return app.successBuilder.Create().
		WithKind(kind).
		WithOutput(retOutput).
		Now()
}

func (app *application) frame(layer layers.Layer, input []byte) (stacks.Frame, error) {
	if input != nil && len(input) <= 0 {
		input = nil
	}

	if input == nil {
		return app.stackFactory.Create().Head(), nil
	}

	assignable, err := app.assignableBuilder.Create().
		WithBytes(input).
		Now()

	if err != nil {
		return nil, err
	}

	inputVariable := layer.Input()
	assignment, err := app.assignmentBuilder.Create().
		WithAssignable(assignable).
		WithName(inputVariable).
		Now()

	if err != nil {
		return nil, err
	}

	assignments, err := app.assignmentsBuilder.Create().
		WithList([]stacks.Assignment{
			assignment,
		}).Now()

	if err != nil {
		return nil, err
	}

	return app.frameBuilder.Create().
		WithAssignments(assignments).
		Now()
}
