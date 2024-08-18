package elements

import (
	"errors"
	"fmt"
)

type elements struct {
	list []Element
	mp   map[string]Element
}

func createElements(
	list []Element,
	mp map[string]Element,
) Elements {
	out := elements{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of element
func (obj *elements) List() []Element {
	return obj.list
}

// Fetch fetches a element by name
func (obj *elements) Fetch(name string) (Element, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the element (name: %s) does not exists", name)
	return nil, errors.New(str)
}
