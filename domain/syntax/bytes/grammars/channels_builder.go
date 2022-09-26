package grammars

import "errors"

type channelsBuilder struct {
	list []Channel
}

func createChannelsBuilder() ChannelsBuilder {
	out := channelsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelsBuilder) Create() ChannelsBuilder {
	return createChannelsBuilder()
}

// WithList adds a list to the builder
func (app *channelsBuilder) WithList(list []Channel) ChannelsBuilder {
	app.list = list
	return app
}

// Now builds a new Channels instance
func (app *channelsBuilder) Now() (Channels, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Channel in order to build a Channels instance")
	}

	return createChannels(app.list), nil
}
