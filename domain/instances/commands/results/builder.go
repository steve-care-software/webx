package results

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	success     Success
	failure     Failure
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		success:     nil,
		failure:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithSuccess adds a success to the builder
func (app *builder) WithSuccess(success Success) Builder {
	app.success = success
	return app
}

// WithFailure adds a failure to the builder
func (app *builder) WithFailure(failure Failure) Builder {
	app.failure = failure
	return app
}

// Now builds a new Result instance
func (app *builder) Now() (Result, error) {
	data := [][]byte{}
	if app.success != nil {
		data = append(data, []byte("isSuccess"))
		data = append(data, app.success.Hash().Bytes())
	}

	if app.failure != nil {
		data = append(data, []byte("isFailure"))
		data = append(data, app.failure.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Result is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.success != nil {
		return createResultWithSuccess(*pHash, app.success), nil
	}

	return createResultWithFailure(*pHash, app.failure), nil
}
