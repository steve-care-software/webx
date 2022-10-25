package selections

type child struct {
	selections Selections
	bytes      []byte
}

func createChildWithSelections(
	selections Selections,
) Child {
	return createChildInternally(selections, nil)
}

func createChildWithBytes(
	bytes []byte,
) Child {
	return createChildInternally(nil, bytes)
}

func createChildInternally(
	selections Selections,
	bytes []byte,
) Child {
	out := child{
		selections: selections,
		bytes:      bytes,
	}

	return &out
}

// IsSelections returns true if there is selections, false otherwise
func (obj *child) IsSelections() bool {
	return obj.selections != nil
}

// Selections returns selections, if any
func (obj *child) Selections() Selections {
	return obj.selections
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *child) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns bytes, if any
func (obj *child) Bytes() []byte {
	return obj.bytes
}
