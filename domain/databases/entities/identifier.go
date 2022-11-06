package entities

type identifier struct {
	section uint
	element uint
}

func createIdentifier(
	section uint,
	element uint,
) Identifier {
	out := identifier{
		section: section,
		element: element,
	}

	return &out
}

// Section returns the section
func (obj *identifier) Section() uint {
	return obj.section
}

// Element returns the element
func (obj *identifier) Element() uint {
	return obj.element
}
