package instructions

type parameter struct {
	element string
	index   uint
	name    string
}

func createParameter(
	element string,
	index uint,
	name string,
) Parameter {
	out := parameter{
		element: element,
		index:   index,
		name:    name,
	}

	return &out
}

// Element returns the element
func (obj *parameter) Element() string {
	return obj.element
}

// Index returns the index
func (obj *parameter) Index() uint {
	return obj.index
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}
