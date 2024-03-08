package assignables

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/executions"
)

type builder struct {
	hashAdapter hash.Adapter
	bytes       bytes.Bytes
	constant    constants.Constant
	execution   executions.Execution
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		constant:    nil,
		execution:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes bytes.Bytes) Builder {
	app.bytes = bytes
	return app
}

// WithConsant adds a constant to the builder
func (app *builder) WithConsant(constant constants.Constant) Builder {
	app.constant = constant
	return app
}

// WithExecution adds an execution to the builder
func (app *builder) WithExecution(execution executions.Execution) Builder {
	app.execution = execution
	return app
}

// Now builds a new Assignable instance
func (app *builder) Now() (Assignable, error) {
	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, []byte("bytes"))
		data = append(data, app.bytes.Hash().Bytes())
	}

	if app.constant != nil {
		data = append(data, []byte("constant"))
		data = append(data, app.constant.Hash().Bytes())
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

	if app.constant != nil {
		return createAssignableWithConstant(*pHash, app.constant), nil
	}

	return createAssignableWithexecution(*pHash, app.execution), nil
}
