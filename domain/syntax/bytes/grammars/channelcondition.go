package grammars

type channelCondition struct {
	prev Token
	next Token
}

func createChannelConditionWithPrevious(
	prev Token,
) ChannelCondition {
	return createChannelConditionInternally(prev, nil)
}

func createChannelConditionWithNext(
	next Token,
) ChannelCondition {
	return createChannelConditionInternally(nil, next)
}

func createChannelConditionWithPreviousAndNext(
	prev Token,
	next Token,
) ChannelCondition {
	return createChannelConditionInternally(prev, next)
}

func createChannelConditionInternally(
	prev Token,
	next Token,
) ChannelCondition {
	out := channelCondition{
		prev: prev,
		next: next,
	}

	return &out
}

// HasPrevious returns true if there is a previous token, false otherwise
func (obj *channelCondition) HasPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous token, if any
func (obj *channelCondition) Previous() Token {
	return obj.prev
}

// HasNext returns true if there is a next token, false otherwise
func (obj *channelCondition) HasNext() bool {
	return obj.next != nil
}

// Next returns the next token, if any
func (obj *channelCondition) Next() Token {
	return obj.next
}
