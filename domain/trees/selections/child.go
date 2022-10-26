package selections

type child struct {
	selections Selections
	content    []byte
}

func createChildWithSelections(
	selections Selections,
) Child {
	return createChildInternally(selections, nil)
}

func createChildWithContent(
	content []byte,
) Child {
	return createChildInternally(nil, content)
}

func createChildInternally(
	selections Selections,
	content []byte,
) Child {
	out := child{
		selections: selections,
		content:    content,
	}

	return &out
}

// Bytes returns the bytes
func (obj *child) Bytes() []byte {
	if obj.IsContent() {
		return obj.Content()
	}

	return obj.Selections().Bytes()
}

// IsSelections returns true if there is selections, false otherwise
func (obj *child) IsSelections() bool {
	return obj.selections != nil
}

// Selections returns selections, if any
func (obj *child) Selections() Selections {
	return obj.selections
}

// IsContent returns true if there is content, false otherwise
func (obj *child) IsContent() bool {
	return obj.content != nil
}

// Content returns content, if any
func (obj *child) Content() []byte {
	return obj.content
}
