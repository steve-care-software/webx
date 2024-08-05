package namespaces

import (
	"errors"
	"fmt"
)

type namespaces struct {
	mp    map[string]Namespace
	list  []Namespace
	names []string
}

func createNamespaces(
	mp map[string]Namespace,
	list []Namespace,
	names []string,
) Namespaces {
	out := namespaces{
		mp:    mp,
		list:  list,
		names: names,
	}

	return &out
}

// List returns the list of namespaces
func (obj *namespaces) List() []Namespace {
	return obj.list
}

// Names returns the names of the namespaces
func (obj *namespaces) Names() []string {
	return obj.names
}

// Fetch fetches a namespace by name
func (obj *namespaces) Fetch(name string) (Namespace, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf(namespaceDoesNotExistsErrPattern, name)
	return nil, errors.New(str)
}

// Index returns the namespace index by name
func (obj *namespaces) Index(name string) (*uint, error) {
	for idx, ins := range obj.list {
		if ins.Name() == name {
			casted := uint(idx)
			return &casted, nil
		}
	}

	str := fmt.Sprintf(namespaceDoesNotExistsErrPattern, name)
	return nil, errors.New(str)
}
