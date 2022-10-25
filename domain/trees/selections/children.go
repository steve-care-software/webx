package selections

type children struct {
	elementName string
	list        []Child
}

func createChildren(
	elementName string,
	list []Child,
) Children {
	out := children{
		elementName: elementName,
		list:        list,
	}

	return &out
}

// ElementName returns the element name
func (obj *children) ElementName() string {
	return obj.elementName
}

// Bytes returns the bytes
func (obj *children) Bytes() []byte {
	content := []byte{}
	for _, oneChild := range obj.list {
		if oneChild.IsBytes() {
			content = append(content, oneChild.Bytes()...)
			continue
		}

		content = append(content, oneChild.Selections().Bytes()...)
	}

	return content
}

// List returns the list of child
func (obj *children) List() []Child {
	return obj.list
}
