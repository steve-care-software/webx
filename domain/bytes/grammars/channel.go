package grammars

type channel struct {
	name      string
	token     Token
	condition ChannelCondition
}

func createChannel(
	name string,
	token Token,
) Channel {
	return createChannelInternally(name, token, nil)
}

func createChannelWithCondition(
	name string,
	token Token,
	condition ChannelCondition,
) Channel {
	return createChannelInternally(name, token, condition)
}

func createChannelInternally(
	name string,
	token Token,
	condition ChannelCondition,
) Channel {
	out := channel{
		name:      name,
		token:     token,
		condition: condition,
	}

	return &out
}

// Name returns the name
func (obj *channel) Name() string {
	return obj.name
}

// Token returns the token
func (obj *channel) Token() Token {
	return obj.token
}

// HasCondition returns true if there is a condition, false otherwise
func (obj *channel) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *channel) Condition() ChannelCondition {
	return obj.condition
}
