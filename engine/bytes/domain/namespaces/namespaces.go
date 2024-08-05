package namespaces

import (
	"errors"
	"fmt"
)

type namespaces struct {
	mp         map[string]Namespace
	list       []Namespace
	activeList []Namespace
	names      []string
	deleted    []string
}

func createNamespaces(
	mp map[string]Namespace,
	list []Namespace,
	activeList []Namespace,
	names []string,
	deleted []string,
) Namespaces {
	out := namespaces{
		mp:         mp,
		list:       list,
		activeList: activeList,
		names:      names,
		deleted:    deleted,
	}

	return &out
}

// List returns the list of namespaces
func (obj *namespaces) List() []Namespace {
	return obj.list
}

// ActiveList returns the list of active namespaces
func (obj *namespaces) ActiveList() []Namespace {
	return obj.activeList
}

// Names returns the names of the namespaces
func (obj *namespaces) Names() []string {
	return obj.names
}

// DeletedNames returns the deleted namespace names
func (obj *namespaces) DeletedNames() []string {
	return obj.deleted
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
