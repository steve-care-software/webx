package instructions

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases"
)

type instructionBuilder struct {
	hashAdapter hash.Adapter
	isStop      bool
	pRaiseError *uint
	condition   Condition
	assignment  assignments.Assignment
	account     accounts.Account
	database    databases.Database
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
		account:     nil,
		database:    nil,
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

// WithAccount adds an account to the builder
func (app *instructionBuilder) WithAccount(account accounts.Account) InstructionBuilder {
	app.account = account
	return app
}

// WithDatabase adds a database to the builder
func (app *instructionBuilder) WithDatabase(database databases.Database) InstructionBuilder {
	app.database = database
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

	if app.account != nil {
		data = append(data, []byte("account"))
		data = append(data, app.account.Hash().Bytes())
	}

	if app.database != nil {
		data = append(data, []byte("database"))
		data = append(data, app.database.Hash().Bytes())
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

	if app.account != nil {
		return createInstructionWithAccount(*pHash, app.account), nil
	}

	if app.database != nil {
		return createInstructionWithDatabase(*pHash, app.database), nil
	}

	return createInstructionWithAssignment(*pHash, app.assignment), nil
}
