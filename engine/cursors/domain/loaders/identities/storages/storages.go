package storages

import (
	"errors"
	"fmt"
)

type storagesIns struct {
	list  []Storage
	mp    map[string]Storage
	names []string
}

func createStorages(
	list []Storage,
	mp map[string]Storage,
	names []string,
) Storages {
	out := storagesIns{
		list:  list,
		mp:    mp,
		names: names,
	}

	return &out
}

// List returns the list
func (obj *storagesIns) List() []Storage {
	return obj.list
}

// Names returns the names
func (obj *storagesIns) Names() []string {
	return obj.names
}

// Fetch fetches a stored identity by name
func (obj *storagesIns) Fetch(name string) (Storage, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the name (%s) could not be found", name)
	return nil, errors.New(str)
}
