package resources

import "errors"

type fieldsBuilder struct {
	list []Field
}

func createFieldsBuilder() FieldsBuilder {
	out := fieldsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the fieldsBuilder
func (app *fieldsBuilder) Create() FieldsBuilder {
	return createFieldsBuilder()
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

	return createFields(app.list), nil
}
