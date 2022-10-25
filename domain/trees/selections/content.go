package selections

type content struct {
	element  Element
	children Children
}

func createContentWithElement(
	element Element,
) Content {
	return createContentInternally(element, nil)
}

func createContentWithChildren(
	children Children,
) Content {
	return createContentInternally(nil, children)
}

func createContentWithElementAndChildren(
	element Element,
	children Children,
) Content {
	return createContentInternally(element, children)
}

func createContentInternally(
	element Element,
	children Children,
) Content {
	out := content{
		element:  element,
		children: children,
	}

	return &out
}

// HasElement returns true if there is an element, false otherwise
func (obj *content) HasElement() bool {
	return obj.element != nil
}

// Element returns the element, if any
func (obj *content) Element() Element {
	return obj.element
}

// HasChildren returns true if there is a children, false otherwise
func (obj *content) HasChildren() bool {
	return obj.children != nil
}

// Children returns the children, if any
func (obj *content) Children() Children {
	return obj.children
}
