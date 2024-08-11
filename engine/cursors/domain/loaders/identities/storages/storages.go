package storages

import (
	"errors"
	"fmt"
)

type storagesIns struct {
	list []Storage
	mp   map[string]Storage
}

func createStorages(
	list []Storage,
	mp map[string]Storage,
) Storages {
	out := storagesIns{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list
func (obj *storagesIns) List() []Storage {
	return obj.list
}

// Fetch fetches a stored identity by name
func (obj *storagesIns) Fetch(name string) (Storage, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the name (%s) could not be found", name)
	return nil, errors.New(str)
}
