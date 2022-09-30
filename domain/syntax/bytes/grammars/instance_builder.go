package grammars

import "errors"

type instanceBuilder struct {
	token      Token
	everything Everything
	isReverse  bool
}

func createInstanceBuilder() InstanceBuilder {
	out := instanceBuilder{
		token:      nil,
		everything: nil,
		isReverse:  false,
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

// IsReverse flags the builder as isReverse
func (app *instanceBuilder) IsReverse() InstanceBuilder {
	app.isReverse = true
	return app
}

// Now builds a new Instance instance
func (app *instanceBuilder) Now() (Instance, error) {
	if app.token != nil {
		content := createInstanceContentWithToken(app.token)
		return createInstance(content, app.isReverse), nil
	}

	if app.everything != nil {
		content := createInstanceContentWithEverything(app.everything)
		return createInstance(content, app.isReverse), nil
	}

	return nil, errors.New("the Instance is invalid")
}
