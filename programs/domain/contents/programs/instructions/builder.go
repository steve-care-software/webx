package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash       *hash.Hash
	pAssignment *hash.Hash
	pExecution  *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:       nil,
		pAssignment: nil,
		pExecution:  nil,
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
func (app *builder) WithAssignment(assignment hash.Hash) Builder {
	app.pAssignment = &assignment
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

	if app.pAssignment != nil {
		content := createContentWithAssignment(app.pAssignment)
		return createInstruction(*app.pHash, content), nil
	}

	if app.pExecution != nil {
		content := createContentWithExecution(app.pExecution)
		return createInstruction(*app.pHash, content), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
