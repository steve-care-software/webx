package resources

import "errors"

type resourceBuilder struct {
	name       string
	key        Field
	fields     Fields
	initialize string
	trigger    string
	children   Resources
}

func createResourceBuilder() ResourceBuilder {
	out := resourceBuilder{
		name:       "",
		key:        nil,
		fields:     nil,
		initialize: "",
		trigger:    "",
		children:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder()
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

// WithInitialize adds an initialize method to the builder
func (app *resourceBuilder) WithInitialize(initialize string) ResourceBuilder {
	app.initialize = initialize
	return app
}

// WithTrigger adds a trigger method to the builder
func (app *resourceBuilder) WithTrigger(trigger string) ResourceBuilder {
	app.trigger = trigger
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

	if app.initialize == "" {
		return nil, errors.New("the initialize method is mandatory in order to build a Resource instance")
	}

	if app.trigger == "" {
		return nil, errors.New("the trigger method is mandatory in order to build a Resource instance")
	}

	if app.children != nil {
		return createResourceWithChildren(
			app.name,
			app.key,
			app.fields,
			app.initialize,
			app.trigger,
			app.children,
		), nil
	}

	return createResource(
		app.name,
		app.key,
		app.fields,
		app.initialize,
		app.trigger,
	), nil
}
