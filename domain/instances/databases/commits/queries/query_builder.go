package queries

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/queries/chunks"
)

type queryBuilder struct {
	hashAdapter hash.Adapter
	bytes       []byte
	chunk       chunks.Chunk
}

func createQueryBuilder(
	hashAdapter hash.Adapter,
) QueryBuilder {
	out := queryBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		chunk:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *queryBuilder) Create() QueryBuilder {
	return createQueryBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *queryBuilder) WithBytes(bytes []byte) QueryBuilder {
	app.bytes = bytes
	return app
}

// WithChunk add chunk to the builder
func (app *queryBuilder) WithChunk(chunk chunks.Chunk) QueryBuilder {
	app.chunk = chunk
	return app
}

// Now builds a new Query instance
func (app *queryBuilder) Now() (Query, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, app.bytes)
	}

	if app.chunk != nil {
		data = append(data, app.chunk.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the bytes or chunk is mandatory in order to build a Query instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createQueryWithBytes(*pHash, app.bytes), nil
	}

	return createQueryWithChunk(*pHash, app.chunk), nil
}
