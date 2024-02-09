package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type assignableBuilder struct {
	hashAdapter hash.Adapter
	bytes       Bytes
	identity    Identity
	execution   Execution
}

func createAssignableBuilder(
	hashAdapter hash.Adapter,
) AssignableBuilder {
	out := assignableBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		identity:    nil,
		execution:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignableBuilder) Create() AssignableBuilder {
	return createAssignableBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *assignableBuilder) WithBytes(bytes Bytes) AssignableBuilder {
	app.bytes = bytes
	return app
}

// WithIdentity add identity to the builder
func (app *assignableBuilder) WithIdentity(identity Identity) AssignableBuilder {
	app.identity = identity
	return app
}

// WithExecution adds an execution to the builder
func (app *assignableBuilder) WithExecution(execution Execution) AssignableBuilder {
	app.execution = execution
	return app
}

// Now builds a new Assignable instance
func (app *assignableBuilder) Now() (Assignable, error) {
	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, []byte("bytes"))
		data = append(data, app.bytes.Hash().Bytes())
	}

	if app.identity != nil {
		data = append(data, []byte("identity"))
		data = append(data, app.identity.Hash().Bytes())
	}

	if app.execution != nil {
		data = append(data, []byte("execution"))
		data = append(data, app.execution.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Assignable is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createAssignableWithBytes(*pHash, app.bytes), nil
	}

	if app.execution != nil {
		return createAssignableWithexecution(*pHash, app.execution), nil
	}

	return createAssignableWithIdentity(*pHash, app.identity), nil
}
