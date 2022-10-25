package selections

type selection struct {
	content Content
}

func createSelection(
	content Content,
) Selection {
	out := selection{
		content: content,
	}

	return &out
}

// ElementName returns the element name
func (obj *selection) ElementName() string {
	if obj.content.HasElement() {
		return obj.content.Element().Value().Grammar().Name()
	}

	return obj.content.Children().ElementName()
}

// Bytes returns the bytes
func (obj *selection) Bytes() []byte {
	if obj.content.HasElement() {
		return obj.content.Element().Bytes()
	}

	return obj.content.Children().Bytes()
}

// Content returns the content
func (obj *selection) Content() Content {
	return obj.content
}
