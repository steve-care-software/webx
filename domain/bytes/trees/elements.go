package trees

type elements struct {
	listWithChannels []Element
	list             []Element
}

func createElements(
	listWithChannels []Element,
	list []Element,
) Elements {
	out := elements{
		listWithChannels: listWithChannels,
		list:             list,
	}

	return &out
}

// List returns the elements
func (obj *elements) List(isChannelsAccepted bool) []Element {
	if isChannelsAccepted {
		return obj.listWithChannels
	}

	return obj.list
}
