package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type listBuilder struct {
	hashAdapter hash.Adapter
	pValue      *uint8
	delimiter   string
}

func createListBuilder(
	hashAdapter hash.Adapter,
) ListBuilder {
	out := listBuilder{
		hashAdapter: hashAdapter,
		pValue:      nil,
		delimiter:   "",
	}

	return &out
}

// Create initializes the builder
func (app *listBuilder) Create() ListBuilder {
	return createListBuilder(
		app.hashAdapter,
	)
}

// WithValue adds a value to the builder
func (app *listBuilder) WithValue(value uint8) ListBuilder {
	app.pValue = &value
	return app
}

// WithDelimiter adds a delimiter to the builder
func (app *listBuilder) WithDelimiter(delimiter string) ListBuilder {
	app.delimiter = delimiter
	return app
}

// Now builds a new List instance
func (app *listBuilder) Now() (List, error) {
	if app.pValue == nil {
		return nil, errors.New("the value is mandatory in order to build a List instance")
	}

	if app.delimiter == "" {
		return nil, errors.New("the delimiter is mandatory in order to build a List instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte{
			*app.pValue,
		},
		[]byte(app.delimiter),
	})

	if err != nil {
		return nil, err
	}

	return createList(*pHash, *app.pValue, app.delimiter), nil
}
