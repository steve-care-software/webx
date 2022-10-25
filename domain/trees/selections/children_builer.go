package selections

import "errors"

type childrenBuilder struct {
	elementName string
	list        []Child
}

func createChildrenBuilder() ChildrenBuilder {
	out := childrenBuilder{
		elementName: "",
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *childrenBuilder) Create() ChildrenBuilder {
	return createChildrenBuilder()
}

// WithElementName adds an element name to the builder
func (app *childrenBuilder) WithElementName(elementName string) ChildrenBuilder {
	app.elementName = elementName
	return app
}

// WithList adds a list to the builder
func (app *childrenBuilder) WithList(list []Child) ChildrenBuilder {
	app.list = list
	return app
}

// Now builds a new Children instance
func (app *childrenBuilder) Now() (Children, error) {
	if app.elementName == "" {
		return nil, errors.New("the element name is mandatory in order to build a Children instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Child in order to build a Children instance")
	}

	return createChildren(app.elementName, app.list), nil
}
