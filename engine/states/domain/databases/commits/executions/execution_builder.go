package executions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions/chunks"
)

type executionBuilder struct {
	hashAdapter hash.Adapter
	bytes       []byte
	chunk       chunks.Chunk
}

func createExecutionBuilder(
	hashAdapter hash.Adapter,
) ExecutionBuilder {
	return &executionBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		chunk:       nil,
	}
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *executionBuilder) WithBytes(bytes []byte) ExecutionBuilder {
	app.bytes = bytes
	return app
}

// WithChunk add chunk to the byuilder
func (app *executionBuilder) WithChunk(chunk chunks.Chunk) ExecutionBuilder {
	app.chunk = chunk
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes != nil && app.chunk != nil {
		return nil, errors.New("only one of bytes or chunk must be provided, not both")
	}

	if app.bytes == nil && app.chunk == nil {
		return nil, errors.New("at least one of bytes or chunk must be provided, none provided")
	}

	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, app.bytes)
	}

	if app.chunk != nil {
		data = append(data, app.chunk.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createExecutionWithBytes(*pHash, app.bytes), nil
	}

	return createExecutionWithChunk(*pHash, app.chunk), nil
}
