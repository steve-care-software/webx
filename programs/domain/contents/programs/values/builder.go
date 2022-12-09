package values

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/contents/programs/assignments"
)

type builder struct {
	pHash      *hash.Hash
	pInput     *uint
	assignment assignments.Assignment
	pExecution *hash.Hash
	pProgram   *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:      nil,
		pInput:     nil,
		assignment: nil,
		pExecution: nil,
		pProgram:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.pHash = &hash
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input uint) Builder {
	app.pInput = &input
	return app
}

// WithAssignment adds an assignment to the builder
func (app *builder) WithAssignment(assignment assignments.Assignment) Builder {
	app.assignment = assignment
	return app
}

// WithExecution adds an execution to the builder
func (app *builder) WithExecution(execution hash.Hash) Builder {
	app.pExecution = &execution
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program hash.Hash) Builder {
	app.pProgram = &program
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Value instance")
	}

	if app.pInput != nil {
		content := createContentWithInput(app.pInput)
		return createValue(*app.pHash, content), nil
	}

	if app.assignment != nil {
		content := createContentWithAssignment(app.assignment)
		return createValue(*app.pHash, content), nil
	}

	if app.pExecution != nil {
		content := createContentWithExecution(app.pExecution)
		return createValue(*app.pHash, content), nil
	}

	if app.pProgram != nil {
		content := createContentWithProgram(app.pProgram)
		return createValue(*app.pHash, content), nil
	}

	return nil, errors.New("the Value is invalid")
}
