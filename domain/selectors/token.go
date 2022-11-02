package selectors

type token struct {
	name        string
	reverseName string
	element     Element
	pContent    *uint
}

func createToken(
	name string,
	reverseName string,
	element Element,
) Token {
	return createTokenInternally(name, reverseName, element, nil)
}

func createTokenWithContentIndex(
	name string,
	reverseName string,
	element Element,
	pContent *uint,
) Token {
	return createTokenInternally(name, reverseName, element, pContent)
}

func createTokenInternally(
	name string,
	reverseName string,
	element Element,
	pContent *uint,
) Token {
	out := token{
		name:        name,
		reverseName: reverseName,
		element:     element,
		pContent:    pContent,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// ReverseName returns the reverse name
func (obj *token) ReverseName() string {
	return obj.reverseName
}

// Element returns the element
func (obj *token) Element() Element {
	return obj.element
}

// HasContent returns true if there is a content index, false otherwise
func (obj *token) HasContent() bool {
	return obj.pContent != nil
}

// Content returns the content index, if any
func (obj *token) Content() *uint {
	return obj.pContent
}
