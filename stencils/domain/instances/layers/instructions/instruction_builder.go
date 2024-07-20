package instructions

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/datastencil/states/domain/hash"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/lists"
)

type instructionBuilder struct {
	hashAdapter hash.Adapter
	isStop      bool
	pRaiseError *uint
	condition   Condition
	assignment  assignments.Assignment
	list        lists.List
	loop        Loop
	execution   executions.Execution
}

func createInstructionBuilder(
	hashAdapter hash.Adapter,
) InstructionBuilder {
	out := instructionBuilder{
		hashAdapter: hashAdapter,
		isStop:      false,
		pRaiseError: nil,
		condition:   nil,
		assignment:  nil,
		list:        nil,
		loop:        nil,
		execution:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder(
		app.hashAdapter,
	)
}

// WithRaiseError raises an error in the builder
func (app *instructionBuilder) WithRaiseError(raiseError uint) InstructionBuilder {
	app.pRaiseError = &raiseError
	return app
}

// WithCondition adds a condition to the builder
func (app *instructionBuilder) WithCondition(condition Condition) InstructionBuilder {
	app.condition = condition
	return app
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment assignments.Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// WithList adds a list to the builder
func (app *instructionBuilder) WithList(list lists.List) InstructionBuilder {
	app.list = list
	return app
}

// WithLoop adds a loop to the builder
func (app *instructionBuilder) WithLoop(loop Loop) InstructionBuilder {
	app.loop = loop
	return app
}

// WithExecution adds an execution to the builder
func (app *instructionBuilder) WithExecution(execution executions.Execution) InstructionBuilder {
	app.execution = execution
	return app
}

// IsStop flags the builder as a stop
func (app *instructionBuilder) IsStop() InstructionBuilder {
	app.isStop = true
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	data := [][]byte{}
	if app.isStop {
		data = append(data, []byte("isStop"))
	}

	if app.pRaiseError != nil {
		data = append(data, []byte("raiseError"))
		data = append(data, []byte(strconv.Itoa(int(*app.pRaiseError))))
	}

	if app.condition != nil {
		data = append(data, []byte("condition"))
		data = append(data, app.condition.Hash().Bytes())
	}

	if app.assignment != nil {
		data = append(data, []byte("assignment"))
		data = append(data, app.assignment.Hash().Bytes())
	}

	if app.list != nil {
		data = append(data, []byte("list"))
		data = append(data, app.list.Hash().Bytes())
	}

	if app.loop != nil {
		data = append(data, []byte("loop"))
		data = append(data, app.loop.Hash().Bytes())
	}

	if app.execution != nil {
		data = append(data, []byte("execution"))
		data = append(data, app.execution.Hash().Bytes())
	}

	dataLength := len(data)
	if dataLength != 1 && dataLength != 2 {
		return nil, errors.New("the Instruction is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isStop {
		return createInstructionWithIsStop(*pHash), nil
	}

	if app.pRaiseError != nil {
		return createInstructionWithRaiseError(*pHash, *app.pRaiseError), nil
	}

	if app.condition != nil {
		return createInstructionWithCondition(*pHash, app.condition), nil
	}

	if app.list != nil {
		return createInstructionWithList(*pHash, app.list), nil
	}

	if app.loop != nil {
		return createInstructionWithLoop(*pHash, app.loop), nil
	}

	if app.execution != nil {
		return createInstructionWithExecution(*pHash, app.execution), nil
	}

	return createInstructionWithAssignment(*pHash, app.assignment), nil
}
