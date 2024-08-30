package resources

import (
	"errors"
	"fmt"
)

type resources struct {
	list []Resource
	mp   map[string]Resource
}

func createResources(
	list []Resource,
	mp map[string]Resource,
) Resources {
	out := resources{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of resources
func (obj *resources) List() []Resource {
	return obj.list
}

// Fetch fetches a resource by name
func (obj *resources) Fetch(name string) (Resource, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the resource (name: %s) could not be found", name)
	return nil, errors.New(str)
}
