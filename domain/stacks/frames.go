package stacks

type frames struct {
	list []Frame
}

func createFrames(
	list []Frame,
) Frames {
	out := frames{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *frames) List() []Frame {
	return obj.list
}
