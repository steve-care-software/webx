package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type channelConditionBuilder struct {
	hashAdapter hash.Adapter
	prev        Token
	next        Token
}

func createChannelConditionBuilder(
	hashAdapter hash.Adapter,
) ChannelConditionBuilder {
	out := channelConditionBuilder{
		hashAdapter: hashAdapter,
		prev:        nil,
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelConditionBuilder) Create() ChannelConditionBuilder {
	return createChannelConditionBuilder(
		app.hashAdapter,
	)
}

// WithPrevious adds a previous token to the builder
func (app *channelConditionBuilder) WithPrevious(previous Token) ChannelConditionBuilder {
	app.prev = previous
	return app
}

// WithPrevious adds a previous token to the builder
func (app *channelConditionBuilder) WithNext(next Token) ChannelConditionBuilder {
	app.next = next
	return app
}

// Now builds a new ChannelCondition instance
func (app *channelConditionBuilder) Now() (ChannelCondition, error) {

	data := [][]byte{}
	if app.next != nil {
		data = append(data, app.next.Hash().Bytes())
	}

	if app.prev != nil {
		data = append(data, app.prev.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the ChannelCondition is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.next != nil && app.prev != nil {
		return createChannelConditionWithPreviousAndNext(*pHash, app.prev, app.next), nil
	}

	if app.next != nil {
		return createChannelConditionWithNext(*pHash, app.next), nil
	}

	return createChannelConditionWithPrevious(*pHash, app.prev), nil
}
