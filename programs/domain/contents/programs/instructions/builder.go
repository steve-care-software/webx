package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/contents/programs/assignments"
)

type builder struct {
	pHash      *hash.Hash
	assignment assignments.Assignment
	pExecution *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:      nil,
		assignment: nil,
		pExecution: nil,
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

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build an Instruction instance")
	}

	if app.assignment != nil {
		content := createContentWithAssignment(app.assignment)
		return createInstruction(*app.pHash, content), nil
	}

	if app.pExecution != nil {
		content := createContentWithExecution(*app.pExecution)
		return createInstruction(*app.pHash, content), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
