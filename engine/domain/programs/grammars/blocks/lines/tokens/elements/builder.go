package elements

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Element
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Element) Builder {
	app.list = list
	return app
}

// Now builds a new Elements instance
func (app *builder) Now() (Elements, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Element in order to build a Elements instance")
	}

	mp := map[string]Element{}
	for _, oneElement := range app.list {
		keyname := oneElement.Name()
		if _, ok := mp[keyname]; ok {
			str := fmt.Sprintf("the Element (name: %s) is a duplicate", keyname)
			return nil, errors.New(str)
		}
		mp[keyname] = oneElement
	}

	return createElements(app.list, mp), nil
}
