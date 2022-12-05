package grammars

import "errors"

type instanceBuilder struct {
	token      Token
	everything Everything
}

func createInstanceBuilder() InstanceBuilder {
	out := instanceBuilder{
		token:      nil,
		everything: nil,
	}

	return &out
}

// Create initializes the builder
func (app *instanceBuilder) Create() InstanceBuilder {
	return createInstanceBuilder()
}

// WithToken adds a token to the builder
func (app *instanceBuilder) WithToken(token Token) InstanceBuilder {
	app.token = token
	return app
}

// WithEverything adds an everything to the builder
func (app *instanceBuilder) WithEverything(everything Everything) InstanceBuilder {
	app.everything = everything
	return app
}

// Now builds a new Instance instance
func (app *instanceBuilder) Now() (Instance, error) {
	if app.token != nil {
		return createInstanceWithToken(app.token), nil
	}

	if app.everything != nil {
		return createInstanceWithEverything(app.everything), nil
	}

	return nil, errors.New("the Instance is invalid")
}
