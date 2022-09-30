package uncovers

type line struct {
	index    uint
	elements []string
}

func createLine(
	index uint,
	elements []string,
) Line {
	out := line{
		index:    index,
		elements: elements,
	}

	return &out
}

// Index returns the index
func (obj *line) Index() uint {
	return obj.index
}

// Elements returns the elements
func (obj *line) Elements() []string {
	return obj.elements
}
