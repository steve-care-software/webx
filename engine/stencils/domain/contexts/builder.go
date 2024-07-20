package contexts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pIdentifier *uint
	head        hash.Hash
	executions  []hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pIdentifier: nil,
		head:        nil,
		executions:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier uint) Builder {
	app.pIdentifier = &identifier
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head hash.Hash) Builder {
	app.head = head
	return app
}

// WithExecutions adds an executions to the builder
func (app *builder) WithExecutions(executions []hash.Hash) Builder {
	app.executions = executions
	return app
}

// Now builds a new Context instance
func (app *builder) Now() (Context, error) {
	if app.pIdentifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Context instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Context instance")
	}

	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Context instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{})
	if err != nil {
		return nil, err
	}

	return createContext(
		*pHash,
		*app.pIdentifier,
		app.head,
		app.executions,
	), nil
}
