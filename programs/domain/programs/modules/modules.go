package modules

import (
	"errors"
	"fmt"
)

type modules struct {
	list []Module
	mp   map[uint]Module
}

func createModules(
	list []Module,
	mp map[uint]Module,
) Modules {
	out := modules{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the modules
func (obj *modules) List() []Module {
	return obj.list
}

// Fetch fetches a module by index
func (obj *modules) Fetch(index uint) (Module, error) {
	if ins, ok := obj.mp[index]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the module (index: %d) is undefined", index)
	return nil, errors.New(str)
}
