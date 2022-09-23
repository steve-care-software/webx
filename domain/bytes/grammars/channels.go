package grammars

type channels struct {
	list []Channel
}

func createChannels(
	list []Channel,
) Channels {
	out := channels{
		list: list,
	}

	return &out
}

// List returns the channels
func (obj *channels) List() []Channel {
	return obj.list
}
