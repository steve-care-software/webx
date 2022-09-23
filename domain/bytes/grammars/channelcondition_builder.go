package grammars

import "errors"

type channelConditionBuilder struct {
	prev Token
	next Token
}

func createChannelConditionBuilder() ChannelConditionBuilder {
	out := channelConditionBuilder{
		prev: nil,
		next: nil,
	}

	return &out
}

// Create initializes the builder
func (app *channelConditionBuilder) Create() ChannelConditionBuilder {
	return createChannelConditionBuilder()
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
	if app.next != nil && app.prev != nil {
		return createChannelConditionWithPreviousAndNext(app.prev, app.next), nil
	}

	if app.next != nil {
		return createChannelConditionWithNext(app.next), nil
	}

	if app.prev != nil {
		return createChannelConditionWithPrevious(app.prev), nil
	}

	return nil, errors.New("the ChannelCondition is invalid")
}
