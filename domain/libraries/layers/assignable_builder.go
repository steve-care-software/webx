package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/executions"
)

type assignableBuilder struct {
	hashAdapter hash.Adapter
	bytes       bytes.Bytes
	execution   executions.Execution
}

func createAssignableBuilder(
	hashAdapter hash.Adapter,
) AssignableBuilder {
	out := assignableBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
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
func (app *assignableBuilder) WithBytes(bytes bytes.Bytes) AssignableBuilder {
	app.bytes = bytes
	return app
}

// WithExecution adds an execution to the builder
func (app *assignableBuilder) WithExecution(execution executions.Execution) AssignableBuilder {
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

	return createAssignableWithexecution(*pHash, app.execution), nil
}
