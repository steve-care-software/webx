package selections

type selection struct {
	elementName string
	list        []Child
}

func createSelection(
	elementName string,
	list []Child,
) Selection {
	out := selection{
		elementName: elementName,
		list:        list,
	}

	return &out
}

// ElementName returns the element name
func (obj *selection) ElementName() string {
	return obj.elementName
}

// Bytes returns the bytes
func (obj *selection) Bytes() []byte {
	content := []byte{}
	for _, oneChild := range obj.list {
		content = append(content, oneChild.Bytes()...)
	}

	return content
}

// List returns the list of child
func (obj *selection) List() []Child {
	return obj.list
}
