package conditions

import (
	"encoding/json"
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type resourceBuilder struct {
	hashAdapter hash.Adapter
	field       Pointer
	value       interface{}
}

func createResourceBuilder(
	hashAdapter hash.Adapter,
) ResourceBuilder {
	out := resourceBuilder{
		hashAdapter: hashAdapter,
		field:       nil,
		value:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder(
		app.hashAdapter,
	)
}

// WithField adds a field to the builder
func (app *resourceBuilder) WithField(field Pointer) ResourceBuilder {
	app.field = field
	return app
}

// WithValue adds a value to the builder
func (app *resourceBuilder) WithValue(value interface{}) ResourceBuilder {
	app.value = value
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
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
