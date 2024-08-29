package instructions

type reference struct {
	element string
	index   uint
}

func createReference(
	element string,
	index uint,
) Reference {
	out := reference{
		element: element,
		index:   index,
	}

	return &out
}

// Element returns the element
func (obj *reference) Element() string {
	return obj.element
}

// Index returns the index
func (obj *reference) Index() uint {
	return obj.index
}
