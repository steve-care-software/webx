package instructions

type elements struct {
	list []Element
}

func createElements(
	list []Element,
) Elements {
	out := elements{
		list: list,
	}

	return &out
}

// List returns the list of element
func (obj *elements) List() []Element {
	return obj.list
}
