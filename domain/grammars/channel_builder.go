package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type channelBuilder struct {
	hashAdapter hash.Adapter
	name        string
	token       Token
	condition   ChannelCondition
}

func createChannelBuilder(
	hashAdapter hash.Adapter,
) ChannelBuilder {
	out := channelBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		token:       nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelBuilder) Create() ChannelBuilder {
	return createChannelBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{
		[]byte(app.name),
		app.token.Hash().Bytes(),
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.condition != nil {
		return createChannelWithCondition(*pHash, app.name, app.token, app.condition), nil
	}

	return createChannel(*pHash, app.name, app.token), nil
}
