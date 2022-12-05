package selectors

type element struct {
	name  string
	index uint
}

func createElement(
	name string,
	index uint,
) Element {
	out := element{
		name:  name,
		index: index,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	return obj.name
}

// Index returns the index
func (obj *element) Index() uint {
	return obj.index
}
