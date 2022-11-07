package databases

type section struct {
	index    uint
	kind     uint8
	pointers Pointers
}

func createSection(
	index uint,
	kind uint8,
) Section {
	return createSectionInternally(index, kind, nil)
}

func createSectionWithPointers(
	index uint,
	kind uint8,
	pointers Pointers,
) Section {
	return createSectionInternally(index, kind, pointers)
}

func createSectionInternally(
	index uint,
	kind uint8,
	pointers Pointers,
) Section {
	out := section{
		index:    index,
		kind:     kind,
		pointers: pointers,
	}

	return &out
}

// Index returns the index
func (obj *section) Index() uint {
	return obj.index
}

// Kind returns the kind
func (obj *section) Kind() uint8 {
	return obj.kind
}

// HasPointers returns true if there is pointers, false otherwise
func (obj *section) HasPointers() bool {
	return obj.pointers != nil
}

// Pointers returns pointers, if any
func (obj *section) Pointers() Pointers {
	return obj.pointers
}
