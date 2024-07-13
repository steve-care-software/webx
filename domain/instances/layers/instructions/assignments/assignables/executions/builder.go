package executions

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	isLength    bool
	pFetch      *uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isLength:    false,
		pFetch:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithFetch adds a fetch to the builder
func (app *builder) WithFetch(fetch uint) Builder {
	app.pFetch = &fetch
	return app
}

// IsLength flags the builder as length
func (app *builder) IsLength() Builder {
	app.isLength = true
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	data := [][]byte{}
	if app.isLength {
		data = append(data, []byte("isLength"))
	}

	if app.pFetch != nil {
		data = append(data, []byte(strconv.Itoa(int(*app.pFetch))))
	}

	if len(data) != 1 {
		return nil, errors.New("the Execution is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pFetch != nil {
		return createExecutionWithFetch(*pHash, *&app.pFetch), nil
	}

	return createExecutionWithLength(*pHash), nil
}
