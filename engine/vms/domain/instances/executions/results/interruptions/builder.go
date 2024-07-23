package interruptions

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/interruptions/failures"
)

type builder struct {
	hashAdapter hash.Adapter
	pStopAtLine *uint
	failure     failures.Failure
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pStopAtLine: nil,
		failure:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithStop adds a stop to the builder
func (app *builder) WithStop(stopLine uint) Builder {
	app.pStopAtLine = &stopLine
	return app
}

// WithFailure adds a stop to the builder
func (app *builder) WithFailure(failure failures.Failure) Builder {
	app.failure = failure
	return app
}

// Now builds a new Interruption instance
func (app *builder) Now() (Interruption, error) {
	data := [][]byte{}
	if app.pStopAtLine != nil {
		data = append(data, []byte("stop"))
		data = append(data, []byte(strconv.Itoa(int(*app.pStopAtLine))))
	}

	if app.failure != nil {
		data = append(data, []byte("failure"))
		data = append(data, app.failure.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Interruption is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pStopAtLine != nil {
		return createInterruptionWithStop(*pHash, app.pStopAtLine), nil
	}

	return createInterruptionWithFailure(*pHash, app.failure), nil
}
