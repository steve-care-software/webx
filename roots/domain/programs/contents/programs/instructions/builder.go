package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pHash      *hash.Hash
	pValue     *hash.Hash
	pExecution *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pHash:      nil,
		pValue:     nil,
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

// WithValue adds an value to the builder
func (app *builder) WithValue(value hash.Hash) Builder {
	app.pValue = &value
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

	if app.pValue != nil {
		content := createContentWithValue(app.pValue)
		return createInstruction(*app.pHash, content), nil
	}

	if app.pExecution != nil {
		content := createContentWithExecution(app.pExecution)
		return createInstruction(*app.pHash, content), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
