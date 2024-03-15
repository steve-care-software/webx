package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type resourceBuilder struct {
	hashAdapter hash.Adapter
	name        string
	key         Field
	fields      Fields
	children    Resources
}

func createResourceBuilder(
	hashAdapter hash.Adapter,
) ResourceBuilder {
	out := resourceBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		key:         nil,
		fields:      nil,
		children:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *resourceBuilder) WithName(name string) ResourceBuilder {
	app.name = name
	return app
}

// WithKey adds a key to the builder
func (app *resourceBuilder) WithKey(key Field) ResourceBuilder {
	app.key = key
	return app
}

// WithFields adds a fields to the builder
func (app *resourceBuilder) WithFields(fields Fields) ResourceBuilder {
	app.fields = fields
	return app
}

// WithChildren adds a children method to the builder
func (app *resourceBuilder) WithChildren(children Resources) ResourceBuilder {
	app.children = children
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Resource instance")
	}

	if app.key == nil {
		return nil, errors.New("the key field is mandatory in order to build a Resource instance")
	}

	if app.fields == nil {
		return nil, errors.New("the fields is mandatory in order to build a Resource instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.key.Hash().Bytes(),
		app.fields.Hash().Bytes(),
	}

	if app.children != nil {
		data = append(data, app.children.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.children != nil {
		return createResourceWithChildren(
			*pHash,
			app.name,
			app.key,
			app.fields,
			app.children,
		), nil
	}

	return createResource(
		*pHash,
		app.name,
		app.key,
		app.fields,
	), nil
}
