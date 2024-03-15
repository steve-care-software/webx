package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type fieldsBuilder struct {
	hashAdapter hash.Adapter
	list        []Field
}

func createFieldsBuilder(
	hashAdapter hash.Adapter,
) FieldsBuilder {
	out := fieldsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the fieldsBuilder
func (app *fieldsBuilder) Create() FieldsBuilder {
	return createFieldsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the fieldsBuilder
func (app *fieldsBuilder) WithList(list []Field) FieldsBuilder {
	app.list = list
	return app
}

// Now builds a new Fields instance
func (app *fieldsBuilder) Now() (Fields, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Field in order to build a Fields instance")
	}

	data := [][]byte{}
	for _, oneField := range app.list {
		data = append(data, oneField.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createFields(*pHash, app.list), nil
}
