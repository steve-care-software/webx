package instructions

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execAccountApp     accounts.Application
	execAssignmentApp  assignments.Application
	execDatabaseApp    databases.Application
	stackBuilder       stacks.Builder
	framesBuilder      stacks.FramesBuilder
	frameBuilder       stacks.FrameBuilder
	assignmentsBuilder stacks.AssignmentsBuilder
	failureBuilder     results.FailureBuilder
}

func createApplication(
	execAccountApp accounts.Application,
	execAssignmentApp assignments.Application,
	execDatabaseApp databases.Application,
	stackBuilder stacks.Builder,
	framesBuilder stacks.FramesBuilder,
	frameBuilder stacks.FrameBuilder,
	assignmentsBuilder stacks.AssignmentsBuilder,
	failureBuilder results.FailureBuilder,
) Application {
	out := application{
		execAccountApp:     execAccountApp,
		execAssignmentApp:  execAssignmentApp,
		execDatabaseApp:    execDatabaseApp,
		stackBuilder:       stackBuilder,
		framesBuilder:      framesBuilder,
		frameBuilder:       frameBuilder,
		assignmentsBuilder: assignmentsBuilder,
		failureBuilder:     failureBuilder,
	}

	return &out
}

// Execute executes the instructions
func (app *application) Execute(stack stacks.Stack, instructions instructions.Instructions) (bool, stacks.Stack, results.Failure, error) {
	currentStack := stack
	list := instructions.List()
	for idx, oneInstruction := range list {
		stop, newStack, failure, err := app.instruction(uint(idx), stack, oneInstruction)
		if err != nil {
			return false, nil, nil, err
		}

		if stop {
			return true, nil, nil, nil
		}

		if failure != nil {
			return false, nil, failure, nil
		}

		currentStack = newStack
	}

	return false, currentStack, nil, nil
}

func (app *application) instruction(idx uint, stack stacks.Stack, instruction instructions.Instruction) (bool, stacks.Stack, results.Failure, error) {
	frame := stack.Head()
	if instruction.IsStop() {
		return true, stack, nil, nil
	}

	if instruction.IsRaiseError() {
		code := instruction.RaiseError()
		failure, err := app.failureBuilder.Create().
			WithIndex(idx).
			WithCode(code).
			IsRaisedInLayer().
			Now()

		if err != nil {
			return false, nil, nil, err
		}

		return false, stack, failure, nil
	}

	if instruction.IsCondition() {
		condition := instruction.Condition()
		variable := condition.Variable()
		value, err := frame.FetchBool(variable)
		if err != nil {
			code := failures.CouldNotFetchConditionFromFrame
			failure, err := app.failureBuilder.Create().
				WithCode(code).
				Now()

			if err != nil {
				return false, nil, nil, err
			}

			return false, stack, failure, nil
		}

		if value {
			instructions := condition.Instructions()
			stop, stack, failure, err := app.Execute(stack, instructions)
			if err != nil {
				return false, nil, nil, err
			}

			if stop {
				return true, nil, nil, nil
			}

			if failure != nil {
				index := failure.Index() + idx
				code := failure.Code()
				builder := app.failureBuilder.Create().WithIndex(index).WithCode(code)
				if failure.IsRaisedInLayer() {
					builder.IsRaisedInLayer()
				}

				failure, err := builder.Now()
				if err != nil {
					return false, nil, nil, err
				}

				return false, stack, failure, nil
			}

			return false, stack, nil, nil
		}

		return true, stack, nil, nil
	}

	if instruction.IsAccount() {
		account := instruction.Account()
		pCode, err := app.execAccountApp.Execute(frame, account)
		if err != nil {
			return false, nil, nil, err
		}

		if pCode != nil {
			failure, err := app.failureBuilder.Create().WithIndex(idx).WithCode(*pCode).Now()
			if err != nil {
				return false, nil, nil, err
			}

			return false, stack, failure, nil
		}

		return false, stack, nil, nil
	}

	if instruction.IsDatabase() {
		database := instruction.Database()
		pCode, err := app.execDatabaseApp.Execute(frame, database)
		if err != nil {
			// log
		}

		if pCode != nil {
			failure, err := app.failureBuilder.Create().WithIndex(idx).WithCode(*pCode).Now()
			if err != nil {
				return false, nil, nil, err
			}

			return false, stack, failure, nil
		}

		return false, stack, nil, nil
	}

	assignment := instruction.Assignment()
	retAssignment, pCode, err := app.execAssignmentApp.Execute(frame, assignment)
	if err != nil {
		return true, nil, nil, err
	}

	if err != nil {
		// log
	}

	if pCode != nil {
		failure, err := app.failureBuilder.Create().WithIndex(idx).WithCode(*pCode).Now()
		if err != nil {
			return false, nil, nil, err
		}

		return false, stack, failure, nil
	}

	assignmentsList := []stacks.Assignment{}
	if frame.HasAssignments() {
		assignmentsList = frame.Assignments().List()
	}

	assignmentsList = append(assignmentsList, retAssignment)
	assignments, err := app.assignmentsBuilder.Create().WithList(assignmentsList).Now()
	if err != nil {
		return true, nil, nil, err
	}

	newFrame, err := app.frameBuilder.Create().WithAssignments(assignments).Now()
	if err != nil {
		return true, nil, nil, err
	}

	framesList := []stacks.Frame{}
	if stack.HasBody() {
		framesList = stack.Body().List()
	}

	framesList = append(framesList, newFrame)
	frames, err := app.framesBuilder.Create().WithList(framesList).Now()
	if err != nil {
		return true, nil, nil, err
	}

	newStack, err := app.stackBuilder.Create().WithFrames(frames).Now()
	if err != nil {
		return true, nil, nil, err
	}

	return false, newStack, nil, nil
}
