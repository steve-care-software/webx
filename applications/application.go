package logics

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/commands"
	"github.com/steve-care-software/datastencil/domain/commands/results"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	layers_bytes "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions"
	links_conditions_resources "github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	hashAdapter        hash.Adapter
	commandsBuilder    commands.Builder
	commandBuilder     commands.CommandBuilder
	resultBuilder      results.Builder
	successBuilder     results.SuccessBuilder
	failureBuilder     results.FailureBuilder
	stackFactory       stacks.Factory
	stackBuilder       stacks.Builder
	framesBuilder      stacks.FramesBuilder
	frameBuilder       stacks.FrameBuilder
	assignmentsBuilder stacks.AssignmentsBuilder
	assignmentBuilder  stacks.AssignmentBuilder
	assignableBuilder  stacks.AssignableBuilder
}

func createApplication(
	hashAdapter hash.Adapter,
	commandsBuilder commands.Builder,
	commandBuilder commands.CommandBuilder,
	resultBuilder results.Builder,
	successBuilder results.SuccessBuilder,
	failureBuilder results.FailureBuilder,
	stackFactory stacks.Factory,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	assignmentBuilder stacks.AssignmentBuilder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		hashAdapter:        hashAdapter,
		commandsBuilder:    commandsBuilder,
		commandBuilder:     commandBuilder,
		resultBuilder:      resultBuilder,
		successBuilder:     successBuilder,
		failureBuilder:     failureBuilder,
		stackFactory:       stackFactory,
		stackBuilder:       stackBuilder,
		framesBuilder:      framesBuilder,
		frameBuilder:       frameBuilder,
		assignmentsBuilder: assignmentsBuilder,
		assignmentBuilder:  assignmentBuilder,
		assignableBuilder:  assignableBuilder,
	}

	return &out
}

// Execute executes the logic application
func (app *application) Execute(input []byte, layer layers.Layer, library libraries.Library, context commands.Commands) (commands.Commands, error) {
	// execute the layer:
	retCommands, err := app.executeLayer(input, layer, library, context)
	if err != nil {
		return nil, err
	}

	// retrieve the link related to our executed layers, if any, from our library:
	if library.HasLinks() {
		// build the list of executed layers:
		layerHashes := []hash.Hash{}
		commandsList := retCommands.List()
		for _, oneCommand := range commandsList {
			layerHashes = append(layerHashes, oneCommand.Layer().Hash())
		}

		link, err := library.Links().FetchByExecutedLayers(layerHashes)
		if err != nil {
			// no link to execute:
			return retCommands, nil
		}

		// execute the link:
		return app.executeLink(
			link,
			library,
			retCommands,
		)
	}

	// execute the link:
	return retCommands, nil
}

func (app *application) executeLink(
	link links.Link,
	library libraries.Library,
	previousCommands commands.Commands,
) (commands.Commands, error) {
	previousResult := previousCommands.Last().Result()
	if previousResult.IsFailure() {
		hash := link.Hash().String()
		str := fmt.Sprintf("the link (hash: %s) cannot execute because the previous result failed", hash)
		return nil, errors.New(str)
	}

	previousSuccess := previousResult.Success()
	currentContext := previousCommands
	elementsList := link.Elements().List()
	for _, oneElement := range elementsList {
		retReceipt, retSuccess, err := app.executeLinkElement(
			oneElement,
			library,
			previousSuccess,
			currentContext,
		)

		if err != nil {
			return nil, err
		}

		currentContext = retReceipt
		previousSuccess = retSuccess
	}

	return currentContext, nil
}

func (app *application) executeLinkElement(
	element elements.Element,
	library libraries.Library,
	previousSuccess results.Success,
	context commands.Commands,
) (commands.Commands, results.Success, error) {
	// execute the layer:
	layerHash := element.Layer()
	layer, err := library.Layers().Fetch(layerHash)
	if err != nil {
		return nil, nil, err
	}

	input := previousSuccess.Bytes()
	retCommands, err := app.Execute(input, layer, library, context)
	if err != nil {
		return nil, nil, err
	}

	retSuccess := app.fetchSuccess(context, retCommands)
	if !element.HasCondition() {
		return retCommands, retSuccess, nil
	}

	result := retCommands.Last().Result()
	if result.IsSuccess() {
		return retCommands, retSuccess, nil
	}

	condition := element.Condition()
	failure := result.Failure()
	if app.matchLinkCondition(condition, failure) {
		return retCommands, previousSuccess, nil
	}

	str := fmt.Sprintf("the layer (hash: %s) did not execute successfully and the link condition did not match", layerHash)
	return nil, nil, errors.New(str)
}

func (app *application) fetchSuccess(
	previous commands.Commands,
	current commands.Commands,
) results.Success {
	result := current.Last().Result()
	if result.IsSuccess() {
		return result.Success()
	}

	return previous.Last().Result().Success()
}

func (app *application) matchLinkCondition(
	condition conditions.Condition,
	failure results.Failure,
) bool {
	resource := condition.Resource()
	isMatch := app.matchLinkConditionResource(resource, failure)
	if isMatch {
		return true
	}

	if !condition.HasNext() {
		return false
	}

	next := condition.Next()
	if next.IsResource() {
		resource := next.Resource()
		return app.matchLinkConditionResource(resource, failure)
	}

	nextCondition := next.Condition()
	return app.matchLinkCondition(nextCondition, failure)
}

func (app *application) matchLinkConditionResource(
	resource links_conditions_resources.Resource,
	failure results.Failure,
) bool {
	actualCode := failure.Code()
	actualIsRaisedInLayer := failure.IsRaisedInLayer()
	expectedCode := resource.Code()
	expectedIsRaisedInLayer := resource.IsRaisedInLayer()
	return actualCode == expectedCode && actualIsRaisedInLayer == expectedIsRaisedInLayer
}

func (app *application) executeLayer(
	input []byte,
	layer layers.Layer,
	library libraries.Library,
	context commands.Commands,
) (commands.Commands, error) {
	assignable, err := app.assignableBuilder.Create().WithBytes(input).Now()
	if err != nil {
		return nil, err
	}

	variable := layer.Input()
	assignment, err := app.assignmentBuilder.Create().WithName(variable).WithAssignable(assignable).Now()
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

	stack, err := app.stackBuilder.Create().WithFrames(frames).Now()
	if err != nil {
		return nil, err
	}

	instructions := layer.Instructions()
	retStack, retFailure, retCommands, err := app.executeInstructions(
		library,
		context,
		layer,
		stack,
		instructions,
	)

	if err != nil {
		return nil, err
	}

	return app.generateCommands(
		input,
		layer,
		retStack,
		retFailure,
		retCommands,
		nil,
	)
}

func (app *application) generateCommands(
	input []byte,
	layer layers.Layer,
	stack stacks.Stack,
	failure results.Failure,
	commands commands.Commands,
	parent commands.Link,
) (commands.Commands, error) {
	commandBuilder := app.commandBuilder.Create().WithInput(input).WithLayer(layer)
	if failure != nil {
		retResult, err := app.resultBuilder.Create().
			WithFailure(failure).
			Now()

		if err != nil {
			return nil, err
		}

		commandBuilder.WithResult(retResult)
	}

	if failure == nil {
		output := layer.Output()
		variable := output.Variable()
		outputBytes, err := stack.Head().FetchBytes(variable)
		if err != nil {
			return nil, err
		}

		if output.HasExecute() {

		}

		kind := output.Kind()
		success, err := app.successBuilder.Create().WithBytes(outputBytes).WithKind(kind).Now()
		if err != nil {
			return nil, err
		}

		retResult, err := app.resultBuilder.Create().
			WithSuccess(success).
			Now()

		if err != nil {
			return nil, err
		}

		commandBuilder.WithResult(retResult)
	}

	if parent != nil {
		commandBuilder.WithParent(parent)
	}

	newCommand, err := commandBuilder.Now()
	if err != nil {
		return nil, err
	}

	commandsList := commands.List()
	commandsList = append(commandsList, newCommand)
	return app.commandsBuilder.Create().
		WithList(commandsList).
		Now()
}

func (app *application) executeInstructions(
	library libraries.Library,
	commands commands.Commands,
	currentLayer layers.Layer,
	stack stacks.Stack,
	instructions layers.Instructions,
) (stacks.Stack, results.Failure, commands.Commands, error) {
	var currentFailure results.Failure
	currentStack := stack
	currentContext := commands
	list := instructions.List()
	for _, oneInstruction := range list {
		stop, retStack, retFailure, retUpdatedCommands, err := app.executeInstruction(
			library,
			currentContext,
			currentLayer,
			currentStack,
			oneInstruction,
		)

		if err != nil {
			return nil, nil, nil, err
		}

		if stop {
			break
		}

		currentStack = retStack
		currentContext = retUpdatedCommands

		if retFailure != nil {
			currentFailure = retFailure
			break
		}
	}

	return currentStack, currentFailure, currentContext, nil
}

func (app *application) executeInstruction(
	library libraries.Library,
	currentContext commands.Commands,
	currentLayer layers.Layer,
	stack stacks.Stack,
	instruction layers.Instruction,
) (bool, stacks.Stack, results.Failure, commands.Commands, error) {
	headFrame := stack.Head()
	if instruction.IsStop() {
		return true, stack, nil, currentContext, nil
	}

	if instruction.IsRaiseError() {
		code := instruction.RaiseError()
		failure, err := app.executeRaiseError(code)
		if err != nil {
			return false, nil, nil, nil, err
		}

		return false, nil, failure, nil, nil
	}

	if instruction.IsCondition() {
		condition := instruction.Condition()
		retStack, retFailure, retCommands, err := app.executeCondition(
			library,
			currentContext,
			currentLayer,
			stack,
			condition,
		)

		if err != nil {
			return false, nil, nil, nil, err
		}

		return false, retStack, retFailure, retCommands, nil
	}

	assignment := instruction.Assignment()
	retFrame, retFailure, retCommands, err := app.executeAssignment(
		library,
		currentContext,
		currentLayer,
		headFrame,
		assignment,
	)

	if err != nil {
		return false, nil, nil, nil, err
	}

	framesList := stack.Frames().List()
	framesList = append(framesList, retFrame)
	updatedFrames, err := app.framesBuilder.Create().
		WithList(framesList).
		Now()

	if err != nil {
		return false, nil, nil, nil, err
	}

	updatedStack, err := app.stackBuilder.Create().
		WithFrames(updatedFrames).
		Now()

	if err != nil {
		return false, nil, nil, nil, err
	}

	return false, updatedStack, retFailure, retCommands, nil
}

func (app *application) executeRaiseError(code uint) (results.Failure, error) {
	failure, err := app.failureBuilder.Create().
		WithCode(code).
		IsRaisedInLayer().
		Now()

	if err != nil {
		return nil, err
	}

	return failure, nil
}

func (app *application) executeCondition(
	library libraries.Library,
	currentContext commands.Commands,
	currentLayer layers.Layer,
	stack stacks.Stack,
	condition layers.Condition,
) (stacks.Stack, results.Failure, commands.Commands, error) {
	variable := condition.Variable()
	boolValue, err := stack.Head().FetchBool(variable)
	if err != nil {

	}

	if boolValue {
		instructions := condition.Instructions()
		return app.executeInstructions(
			library,
			currentContext,
			currentLayer,
			stack,
			instructions,
		)
	}

	return stack, nil, currentContext, nil
}

func (app *application) executeAssignment(
	library libraries.Library,
	currentContext commands.Commands,
	currentLayer layers.Layer,
	frame stacks.Frame,
	assignment layers.Assignment,
) (stacks.Frame, results.Failure, commands.Commands, error) {
	assignable := assignment.Assignable()
	retAssignable, failure, receipts, err := app.executeAssignable(
		library,
		currentContext,
		currentLayer,
		frame,
		assignable,
	)

	if err != nil {
		return nil, nil, nil, err
	}

	name := assignment.Name()
	currentAssignmentsList := []stacks.Assignment{}
	if frame.HasAssignments() {
		currentAssignmentsList = frame.Assignments().List()
	}

	newAssignment, err := app.assignmentBuilder.Create().
		WithName(name).
		WithAssignable(retAssignable).
		Now()

	if err != nil {
		return nil, nil, nil, err
	}

	currentAssignmentsList = append(currentAssignmentsList, newAssignment)
	assignments, err := app.assignmentsBuilder.Create().
		WithList(currentAssignmentsList).
		Now()

	if err != nil {
		return nil, nil, nil, err
	}

	retFrame, err := app.frameBuilder.Create().
		WithAssignments(assignments).
		Now()

	if err != nil {
		return nil, nil, nil, err
	}

	return retFrame, failure, receipts, nil
}

func (app *application) executeAssignable(
	library libraries.Library,
	currentContext commands.Commands,
	currentLayer layers.Layer,
	frame stacks.Frame,
	assignable layers.Assignable,
) (stacks.Assignable, results.Failure, commands.Commands, error) {
	if assignable.IsBytes() {
		bytesIns := assignable.Bytes()
		retAssignable, err := app.executeBytes(
			frame,
			bytesIns,
		)

		if err != nil {
			return nil, nil, nil, err
		}

		return retAssignable, nil, currentContext, nil
	}

	execution := assignable.Execution()
	return app.executeExecution(
		library,
		currentContext,
		currentLayer,
		frame,
		execution,
	)
}

func (app *application) executeExecution(
	library libraries.Library,
	currentContext commands.Commands,
	currentLayer layers.Layer,
	frame stacks.Frame,
	execution layers.Execution,
) (stacks.Assignable, results.Failure, commands.Commands, error) {
	inputVariable := execution.Input()
	input, err := frame.FetchBytes(inputVariable)
	if err != nil {
		return nil, nil, nil, err
	}

	layerToExecute := currentLayer
	if execution.HasLayer() {
		layerHashVariable := execution.Layer()
		layerHash, err := frame.FetchHash(layerHashVariable)
		if err != nil {
			return nil, nil, nil, err
		}

		layer, err := library.Layers().Fetch(layerHash)
		if err != nil {
			return nil, nil, nil, err
		}

		layerToExecute = layer
	}

	retCommands, err := app.Execute(
		input,
		layerToExecute,
		library,
		currentContext,
	)

	if err != nil {
		return nil, nil, nil, err
	}

	result := retCommands.Last().Result()
	if result.IsSuccess() {
		data := result.Success().Bytes()
		retAssignable, err := app.assignableBuilder.Create().
			WithBytes(data).
			Now()

		if err != nil {
			return nil, nil, nil, err
		}

		return retAssignable, nil, retCommands, nil
	}

	failure := result.Failure()
	return nil, failure, retCommands, nil
}

func (app *application) executeBytes(frame stacks.Frame, bytesIns layers_bytes.Bytes) (stacks.Assignable, error) {
	if bytesIns.IsJoin() {
		variables := bytesIns.Join()
		return app.executeJoin(frame, variables)
	}

	if bytesIns.IsCompare() {
		variables := bytesIns.Join()
		return app.executeCompare(frame, variables)
	}

	variable := bytesIns.HashBytes()
	return app.executeHashBytes(frame, variable)
}

func (app *application) executeJoin(frame stacks.Frame, variables []string) (stacks.Assignable, error) {
	output := []byte{}
	for _, oneVariable := range variables {
		data, err := frame.FetchBytes(oneVariable)
		if err != nil {
			return nil, err
		}

		output = append(output, data...)
	}

	return app.assignableBuilder.Create().
		WithBytes(output).
		Now()
}

func (app *application) executeCompare(frame stacks.Frame, variables []string) (stacks.Assignable, error) {
	boolValue := true
	var lastBytes []byte
	for _, oneVariable := range variables {
		data, err := frame.FetchBytes(oneVariable)
		if err != nil {
			return nil, err
		}

		if lastBytes == nil {
			lastBytes = data
			continue
		}

		if !bytes.Equal(lastBytes, data) {
			boolValue = false
			break
		}
	}

	return app.assignableBuilder.Create().
		WithBool(boolValue).
		Now()
}

func (app *application) executeHashBytes(frame stacks.Frame, variable string) (stacks.Assignable, error) {
	data, err := frame.FetchBytes(variable)
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithHash(*pHash).
		Now()
}
