package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type channelsBuilder struct {
	hashAdapter hash.Adapter
	list        []Channel
}

func createChannelsBuilder(
	hashAdapter hash.Adapter,
) ChannelsBuilder {
	out := channelsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelsBuilder) Create() ChannelsBuilder {
	return createChannelsBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{}
	for _, oneChannel := range app.list {
		data = append(data, oneChannel.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createChannels(*pHash, app.list), nil
}
