package resources

import (
	"encoding/json"
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
)

type builder struct {
	hashAdapter hash.Adapter
	field       pointers.Pointer
	value       interface{}
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		field:       nil,
		value:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithField adds a field to the builder
func (app *builder) WithField(field pointers.Pointer) Builder {
	app.field = field
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value interface{}) Builder {
	app.value = value
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	data := [][]byte{}
	if app.field != nil {
		data = append(data, app.field.Hash().Bytes())
	}

	if app.value != nil {
		bytes, err := json.Marshal(app.value)
		if err != nil {
			return nil, err
		}

		data = append(data, bytes)
	}

	if len(data) != 1 {
		return nil, errors.New("the Resource is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.field != nil {
		return createResourceWithField(*pHash, app.field), nil
	}

	return createResourceWithValue(*pHash, app.value), nil
}
