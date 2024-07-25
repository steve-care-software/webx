package instructions

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/executions"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/lists"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/interruptions/failures"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

type application struct {
	assignmentApp       assignments.Application
	listApp             lists.Application
	executionApp        executions.Application
	stackBuilder        stacks.Builder
	framesBuilder       stacks.FramesBuilder
	frameBuilder        stacks.FrameBuilder
	assignmentsBuilder  stacks.AssignmentsBuilder
	interruptionBuilder interruptions.Builder
	failureBuilder      failures.Builder
}

func createApplication(
	assignmentApp assignments.Application,
	listApp lists.Application,
	executionApp executions.Application,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	interruptionBuilder interruptions.Builder,
	failureBuilder failures.Builder,
) Application {
	out := application{
		assignmentApp:       assignmentApp,
		listApp:             listApp,
		executionApp:        executionApp,
		stackBuilder:        stackBuilder,
		framesBuilder:       framesBuilder,
		frameBuilder:        frameBuilder,
		assignmentsBuilder:  assignmentsBuilder,
		interruptionBuilder: interruptionBuilder,
		failureBuilder:      failureBuilder,
	}

	return &out
}

// Execute executes instructions with the provided stack
func (app *application) Execute(
	frame stacks.Frame,
	instructions instructions.Instructions,
) (stacks.Frame, interruptions.Interruption, error) {
	currentFrame := frame
	list := instructions.List()
	for idx, onInstruction := range list {
		castedIndex := uint(idx)
		retFrame, retInterruption, err := app.instruction(castedIndex, currentFrame, onInstruction)
		if err != nil {
			return nil, nil, err
		}

		if retInterruption != nil {
			return nil, retInterruption, nil
		}

		currentFrame = retFrame
	}

	return currentFrame, nil, nil
}

func (app *application) instruction(
	index uint,
	frame stacks.Frame,
	instruction instructions.Instruction,
) (stacks.Frame, interruptions.Interruption, error) {
	if instruction.IsStop() {
		interruption, err := app.interrupt(&index, nil)
		if err != nil {
			return nil, nil, err
		}

		return nil, interruption, nil
	}

	if instruction.IsRaiseError() {
		code := instruction.RaiseError()
		interruption, err := app.interrupt(nil, &failure{
			index:         index,
			code:          code,
			raisedInLayer: true,
			message:       "Error raised in layer",
		})

		if err != nil {
			return nil, nil, err
		}

		return nil, interruption, nil
	}

	if instruction.IsCondition() {
		condition := instruction.Condition()
		variable := condition.Variable()
		value, err := frame.FetchBool(variable)
		if err != nil {
			return nil, nil, err
		}

		if value {
			instructions := condition.Instructions()
			retStack, retInterruption, err := app.Execute(frame, instructions)
			if err != nil {
				return nil, nil, err
			}

			if retInterruption != nil {
				return nil, retInterruption, nil
			}

			return retStack, nil, nil
		}

		return frame, nil, nil
	}

	if instruction.IsAssignment() {
		assignment := instruction.Assignment()
		retAssignment, pCode, err := app.assignmentApp.Execute(frame, assignment)
		if pCode != nil {
			message := ""
			if err != nil {
				message = err.Error()
			}

			interruption, err := app.interrupt(nil, &failure{
				index:         index,
				code:          *pCode,
				raisedInLayer: false,
				message:       message,
			})

			if err != nil {
				return nil, nil, err
			}

			return nil, interruption, nil
		}

		if err != nil {
			return nil, nil, err
		}

		retFrame, err := app.frame(frame, retAssignment)
		if err != nil {
			return nil, nil, err
		}

		return retFrame, nil, nil
	}

	if instruction.IsList() {
		list := instruction.List()
		retAssignment, pCode, err := app.listApp.Execute(frame, list)
		if pCode != nil {
			message := ""
			if err != nil {
				message = err.Error()
			}

			interruption, err := app.interrupt(nil, &failure{
				index:         index,
				code:          *pCode,
				raisedInLayer: false,
				message:       message,
			})

			if err != nil {
				return nil, nil, err
			}

			return nil, interruption, nil
		}

		if err != nil {
			return nil, nil, err
		}

		retFrame, err := app.frame(frame, retAssignment)
		if err != nil {
			return nil, nil, err
		}

		return retFrame, nil, nil
	}

	loop := instruction.Loop()
	amoountVar := loop.Amount()
	pAmount, err := frame.FetchUnsignedInt(amoountVar)
	if err != nil {
		return nil, nil, err
	}

	currentFrame := frame
	casted := int(*pAmount)
	instructions := loop.Instructions()
	for i := 0; i < casted; i++ {
		retFrame, retInterruption, err := app.Execute(frame, instructions)
		if err != nil {
			return nil, nil, err
		}

		if retInterruption != nil {
			return nil, retInterruption, nil
		}

		currentFrame = retFrame
	}

	return currentFrame, nil, nil
}

func (app *application) frame(
	currentFrame stacks.Frame,
	newAssignment stacks.Assignment,
) (stacks.Frame, error) {
	assignmentsList := []stacks.Assignment{}
	if currentFrame.HasAssignments() {
		assignmentsList = currentFrame.Assignments().List()
	}

	assignmentsList = append(assignmentsList, newAssignment)
	assignments, err := app.assignmentsBuilder.Create().
		WithList(assignmentsList).
		Now()

	if err != nil {
		return nil, err
	}

	frame, err := app.frameBuilder.Create().
		WithAssignments(assignments).
		Now()

	if err != nil {
		return nil, err
	}

	return frame, nil
}

func (app *application) interrupt(
	pStopIndex *uint,
	pFailure *failure,
) (interruptions.Interruption, error) {
	builder := app.interruptionBuilder.Create()
	if pStopIndex != nil {
		builder.WithStop(*pStopIndex)
	}

	if pFailure != nil {
		failure, err := app.fail(*pFailure)
		if err != nil {
			return nil, err
		}

		builder.WithFailure(failure)
	}

	return builder.Now()
}

func (app *application) fail(
	failure failure,
) (failures.Failure, error) {
	builder := app.failureBuilder.Create().
		WithIndex(failure.index).
		WithCode(failure.code)

	if failure.message != "" {
		builder.WithMessage(failure.message)
	}

	if failure.raisedInLayer {
		builder.IsRaisedInLayer()
	}

	return builder.Now()
}
