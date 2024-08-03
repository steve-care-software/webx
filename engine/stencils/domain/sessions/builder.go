package sessions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
)

type builder struct {
	hashAdapter hash.Adapter
	hash        hash.Hash
	executions  []executions.Executions
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		hash:        nil,
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

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.hash = hash
	return app
}

// WithExecutions add executions to the builder
func (app *builder) WithExecutions(executions []executions.Executions) Builder {
	app.executions = executions
	return app
}

// Now builds a new Session instance
func (app *builder) Now() (Session, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Session instance")
	}

	if app.executions != nil && len(app.executions) <= 0 {
		app.executions = nil
	}

	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Session instance")
	}

	bytes := [][]byte{}
	for _, oneExecutions := range app.executions {
		bytes = append(bytes, oneExecutions.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	return createSession(
		app.hash,
		app.executions,
		*pHash,
	), nil
}
