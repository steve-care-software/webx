package logics

import (
	"github.com/steve-care-software/datastencil/applications/logics/binaries"
	"github.com/steve-care-software/datastencil/applications/logics/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	execution_links "github.com/steve-care-software/datastencil/domain/instances/executions/links"
	execution_layers "github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/success"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics"
	bridged_layers "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/references"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	instructionsApplication    instructions.Application
	binaryApplication          binaries.Application
	executionLayersBuilder     execution_layers.Builder
	executionLayerBuilder      execution_layers.LayerBuilder
	executionLinkBuilder       execution_links.Builder
	resultBuilder              results.Builder
	resultSuccessBuilder       success.Builder
	resultSuccessOutputBuilder outputs.Builder
	stackBuilder               stacks.Builder
	framesBuilder              stacks.FramesBuilder
	frameBuilder               stacks.FrameBuilder
	assignmentsBuilder         stacks.AssignmentsBuilder
	assignmentBuilder          stacks.AssignmentBuilder
	assignablesBuilder         stacks.AssignablesBuilder
	assignableBuilder          stacks.AssignableBuilder
}

func createApplication(
	instructionsApplication instructions.Application,
	binaryApplication binaries.Application,
	executionLayersBuilder execution_layers.Builder,
	executionLayerBuilder execution_layers.LayerBuilder,
	executionLinkBuilder execution_links.Builder,
	resultBuilder results.Builder,
	resultSuccessBuilder success.Builder,
	resultSuccessOutputBuilder outputs.Builder,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	assignmentBuilder stacks.AssignmentBuilder,
	assignablesBuilder stacks.AssignablesBuilder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		instructionsApplication:    instructionsApplication,
		binaryApplication:          binaryApplication,
		executionLayersBuilder:     executionLayersBuilder,
		executionLayerBuilder:      executionLayerBuilder,
		executionLinkBuilder:       executionLinkBuilder,
		resultBuilder:              resultBuilder,
		resultSuccessBuilder:       resultSuccessBuilder,
		resultSuccessOutputBuilder: resultSuccessOutputBuilder,
		stackBuilder:               stackBuilder,
		framesBuilder:              framesBuilder,
		frameBuilder:               frameBuilder,
		assignmentsBuilder:         assignmentsBuilder,
		assignmentBuilder:          assignmentBuilder,
		assignablesBuilder:         assignablesBuilder,
		assignableBuilder:          assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte, logic logics.Logic) (execution_links.Link, error) {
	return app.execute(input, logic, nil)
}

// Execute executes the application with context
func (app *application) ExecuteWithContext(input []byte, logic logics.Logic, context executions.Executions) (execution_links.Link, error) {
	return app.execute(input, logic, context)
}

func (app *application) execute(input []byte, logic logics.Logic, context executions.Executions) (execution_links.Link, error) {
	references := logic.References()
	bridges := logic.Bridges()
	executedLayersList := []execution_layers.Layer{}
	link := logic.Link()
	elementsList := link.Elements().List()
	for _, oneElement := range elementsList {
		layerPath := oneElement.Layer()
		bridge, err := bridges.Fetch(layerPath)
		if err != nil {
			return nil, err
		}

		layer := bridge.Layer()
		retResult, err := app.executeLayer(input, layer, references)
		if err != nil {
			return nil, err
		}

		executedLayer, err := app.executionLayerBuilder.Create().WithInput(input).WithSource(layer).WithResult(retResult).Now()
		if err != nil {
			return nil, err
		}

		executedLayersList = append(executedLayersList, executedLayer)
		if retResult.IsSuccess() {
			continue
		}

		interruption := retResult.Interruption()
		if interruption.IsStop() {
			break
		}

		failure := interruption.Failure()
		code := failure.Code()
		isRaisedInLayer := failure.IsRaisedInLayer()
		condition := oneElement.Condition()
		if !app.respectCondition(condition, code, isRaisedInLayer) {
			break
		}
	}

	executedLayers, err := app.executionLayersBuilder.Create().
		WithList(executedLayersList).
		Now()

	if err != nil {
		return nil, err
	}

	return app.executionLinkBuilder.Create().
		WithInput(input).
		WithLayers(executedLayers).
		WithSource(link).
		Now()
}

func (app *application) executeLayer(
	input []byte,
	layer bridged_layers.Layer,
	references references.References,
) (results.Result, error) {
	stack, err := app.initStack(input, layer, references)
	if err != nil {
		return nil, err
	}

	instructions := layer.Instructions()
	retStack, retInterruption, err := app.instructionsApplication.Execute(stack, instructions)
	if err != nil {
		return nil, err
	}

	builder := app.resultBuilder.Create()
	if retInterruption != nil {
		builder.WithInterruption(retInterruption)
	}

	if retStack != nil {
		layerOutput := layer.Output()
		kind := layerOutput.Kind()
		variable := layerOutput.Variable()
		value, err := retStack.Head().FetchBytes(variable)
		if err != nil {
			return nil, err
		}

		outputBuilder := app.resultSuccessOutputBuilder.Create().WithInput(value)
		if layerOutput.HasExecute() {
			executeCmd := layerOutput.Execute()
			retOutput, err := app.binaryApplication.Execute(value, executeCmd)
			if err != nil {
				return nil, err
			}

			outputBuilder.WithExecute(retOutput)
		}

		output, err := outputBuilder.Now()
		if err != nil {
			return nil, err
		}

		success, err := app.resultSuccessBuilder.Create().
			WithKind(kind).
			WithOutput(output).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithSuccess(success)
	}

	return builder.Now()
}

func (app *application) initStack(
	input []byte,
	layer bridged_layers.Layer,
	references references.References,
) (stacks.Stack, error) {
	inputAssignment, err := app.initInputAssignment(input, layer)
	if err != nil {
		return nil, err
	}

	assignmentsList := []stacks.Assignment{
		inputAssignment,
	}

	referencesList := references.List()
	for _, oneReference := range referencesList {
		instance := oneReference.Instance()
		assignable, err := app.assignableBuilder.Create().WithInstance(instance).Now()
		if err != nil {
			return nil, err
		}

		name := oneReference.Variable()
		assignment, err := app.assignmentBuilder.Create().
			WithAssignable(assignable).
			WithName(name).
			Now()

		if err != nil {
			return nil, err
		}

		assignmentsList = append(assignmentsList, assignment)
	}

	assignments, err := app.assignmentsBuilder.Create().WithList(assignmentsList).Now()
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

	return app.stackBuilder.Create().
		WithFrames(frames).
		Now()
}

func (app *application) initInputAssignment(
	input []byte,
	layer bridged_layers.Layer,
) (stacks.Assignment, error) {
	assignable, err := app.assignableBuilder.Create().WithBytes(input).Now()
	if err != nil {
		return nil, err
	}

	variable := layer.Input()
	return app.assignmentBuilder.Create().
		WithAssignable(assignable).
		WithName(variable).
		Now()
}

func (app *application) respectCondition(
	exepectedCondition conditions.Condition,
	code uint,
	isRaisedInLayer bool,
) bool {
	resource := exepectedCondition.Resource()
	expectedCode := resource.Code()
	expectedIsRaisedInLayer := resource.IsRaisedInLayer()
	if expectedCode == code && expectedIsRaisedInLayer == isRaisedInLayer {
		return true
	}

	if !exepectedCondition.HasNext() {
		return false
	}

	next := exepectedCondition.Next()
	return app.respectCondition(next, code, isRaisedInLayer)
}
