package coverages

type line struct {
	list []Element
}

func createLine(
	list []Element,
) Line {
	out := line{
		list: list,
	}

	return &out
}

// List returns the elements
func (obj *line) List() []Element {
	return obj.list
}
