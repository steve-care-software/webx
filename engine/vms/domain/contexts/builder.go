package contexts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
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

	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Context instance")
	}

	data := [][]byte{
		[]byte(fmt.Sprintf("%d", *app.pIdentifier)),
	}

	for _, oneHash := range app.executions {
		data = append(data, oneHash.Bytes())
	}

	if app.head != nil {
		data = append(data, app.head.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.head != nil {
		return createContextWithHead(*pHash, *app.pIdentifier, app.executions, app.head), nil
	}

	return createContext(
		*pHash,
		*app.pIdentifier,
		app.executions,
	), nil
}
