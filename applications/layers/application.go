package layers

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"github.com/steve-care-software/datastencil/applications/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execInsApp         instructions.Application
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
func (app *application) Execute(input []byte, layer layers.Layer, stack stacks.Stack) (results.Result, error) {
	inName := layer.Input()
	retStack, err := app.newStack(input, inName, stack)
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

func (app *application) newStack(input []byte, variable string, context stacks.Stack) (stacks.Stack, error) {
	currentFrames := []stacks.Frame{}
	if context != nil {
		currentFrames = context.Frames().List()
	}

	assignable, err := app.assignableBuilder.Create().
		WithBytes(input).
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

	assignments, err := app.assignmentsBuilder.Create().
		WithList([]stacks.Assignment{
			assignment,
		}).Now()

	if err != nil {
		return nil, err
	}

	newFrame, err := app.frameBuilder.Create().
		WithAssignments(assignments).
		Now()

	if err != nil {
		return nil, err
	}

	currentFrames = append(currentFrames, newFrame)
	frames, err := app.framesBuilder.Create().
		WithList(currentFrames).
		Now()

	if err != nil {
		return nil, err
	}

	return app.stackBuilder.Create().
		WithFrames(frames).
		Now()
}
