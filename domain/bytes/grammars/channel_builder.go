package grammars

import "errors"

type channelBuilder struct {
	name      string
	token     Token
	condition ChannelCondition
}

func createChannelBuilder() ChannelBuilder {
	out := channelBuilder{
		name:      "",
		token:     nil,
		condition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelBuilder) Create() ChannelBuilder {
	return createChannelBuilder()
}

// WithName adds a name to the builder
func (app *channelBuilder) WithName(name string) ChannelBuilder {
	app.name = name
	return app
}

// WithToken adds a token to the builder
func (app *channelBuilder) WithToken(token Token) ChannelBuilder {
	app.token = token
	return app
}

// WithCondition adds a condition to the builder
func (app *channelBuilder) WithCondition(condition ChannelCondition) ChannelBuilder {
	app.condition = condition
	return app
}

// Now builds a new Channel instance
func (app *channelBuilder) Now() (Channel, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Channel instance")
	}

	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Channel instance")
	}

	if app.condition != nil {
		return createChannelWithCondition(app.name, app.token, app.condition), nil
	}

	return createChannel(app.name, app.token), nil
}
