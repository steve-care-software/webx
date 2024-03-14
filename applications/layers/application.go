package layers

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/steve-care-software/datastencil/applications/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/commands"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execInsApp         instructions.Application
	stackFactory       stacks.Factory
	stackBuilder       stacks.Builder
	framesBuilder      stacks.FramesBuilder
	frameBuilder       stacks.FrameBuilder
	assignmentsBuilder stacks.AssignmentsBuilder
	assignmentBuilder  stacks.AssignmentBuilder
	assignableBuilder  stacks.AssignableBuilder
	resultBuilder      results.Builder
	successBuilder     results.SuccessBuilder
	outputBuilder      results.OutputBuilder
	tempBaseDir        string
}

func createApplication(
	execInsApp instructions.Application,
	stackFactory stacks.Factory,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	assignmentBuilder stacks.AssignmentBuilder,
	assignableBuilder stacks.AssignableBuilder,
	resultBuilder results.Builder,
	successBuilder results.SuccessBuilder,
	outputBuilder results.OutputBuilder,
	tempBaseDir string,
) Application {
	out := application{
		execInsApp:         execInsApp,
		stackFactory:       stackFactory,
		stackBuilder:       stackBuilder,
		framesBuilder:      framesBuilder,
		frameBuilder:       frameBuilder,
		assignmentsBuilder: assignmentsBuilder,
		assignmentBuilder:  assignmentBuilder,
		assignableBuilder:  assignableBuilder,
		resultBuilder:      resultBuilder,
		successBuilder:     successBuilder,
		outputBuilder:      outputBuilder,
		tempBaseDir:        tempBaseDir,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte, layer layers.Layer, context commands.Commands) (results.Result, error) {
	previousFrame, err := app.commandsToFrame(context)
	if err != nil {
		return nil, err
	}

	inName := layer.Input()
	retStack, err := app.newStack(input, inName, previousFrame)
	if err != nil {
		return nil, err
	}

	instructions := layer.Instructions()
	retStack, retInterruption, err := app.execInsApp.Execute(retStack, instructions)
	if err != nil {
		return nil, err
	}

	if retInterruption != nil {
		return app.resultBuilder.Create().
			WithInterruption(retInterruption).
			Now()
	}

	output := layer.Output()
	outName := output.Variable()
	retBytes, err := retStack.Head().FetchBytes(outName)
	if err != nil {
		return nil, err
	}

	kind := output.Kind()
	outputBuilder := app.outputBuilder.Create().WithInput(retBytes)
	if output.HasExecute() {
		command := output.Execute()
		bytesAfterExecute, err := app.executeThenReturn(retBytes, command)
		if err != nil {
			return nil, err
		}

		outputBuilder.WithExecute(bytesAfterExecute)
	}

	outputIns, err := outputBuilder.Now()
	if err != nil {
		return nil, err
	}

	success, err := app.successBuilder.Create().
		WithOutput(outputIns).
		WithKind(kind).
		Now()

	if err != nil {
		return nil, err
	}

	return app.resultBuilder.Create().
		WithSuccess(success).
		Now()
}

func (app *application) executeThenReturn(value []byte, args []string) ([]byte, error) {
	pFile, err := ioutil.TempFile(app.tempBaseDir, "*")
	if err != nil {
		return nil, err
	}

	path := filepath.Join(
		app.tempBaseDir,
		pFile.Name(),
	)

	return exec.Command(path, args...).Output()
}

func (app *application) newStack(input []byte, variable string, previousFrame stacks.Frame) (stacks.Stack, error) {
	assignments, err := app.buildAssignmentsWithByteVariables(map[string][]byte{
		variable: input,
	})

	if err != nil {
		return nil, err
	}

	currentFrames := []stacks.Frame{}
	if previousFrame != nil {
		currentFrames = append(currentFrames, previousFrame)
	}

	return app.buildStackWithAssignmentsAndPreviousFrames(
		assignments,
		currentFrames,
	)
}

func (app *application) commandsToFrame(commands commands.Commands) (stacks.Frame, error) {
	if commands == nil {
		return nil, nil
	}

	variables := map[string][]byte{}
	list := commands.List()
	for _, oneCommand := range list {
		input := oneCommand.Input()
		variable := oneCommand.Layer().Input()
		variables[variable] = input
	}

	assignments, err := app.buildAssignmentsWithByteVariables(variables)
	if err != nil {
		return nil, err
	}

	return app.frameBuilder.Create().
		WithAssignments(assignments).
		Now()
}

func (app *application) buildStackWithAssignmentsAndPreviousFrames(
	assignments stacks.Assignments,
	previous []stacks.Frame,
) (stacks.Stack, error) {
	newFrame, err := app.frameBuilder.Create().
		WithAssignments(assignments).
		Now()

	if err != nil {
		return nil, err
	}

	previous = append(previous, newFrame)
	frames, err := app.framesBuilder.Create().
		WithList(previous).
		Now()

	if err != nil {
		return nil, err
	}

	return app.stackBuilder.Create().
		WithFrames(frames).
		Now()
}

func (app *application) buildAssignmentsWithByteVariables(input map[string][]byte) (stacks.Assignments, error) {
	list := []stacks.Assignment{}
	for variable, data := range input {
		assignable, err := app.assignableBuilder.Create().
			WithBytes(data).
			Now()

		if err != nil {
			return nil, err
		}

		assignment, err := app.assignmentBuilder.Create().
			WithName(variable).
			WithAssignable(assignable).
			Now()

		if err != nil {
			return nil, err
		}

		list = append(list, assignment)
	}

	return app.assignmentsBuilder.Create().
		WithList(list).
		Now()
}
