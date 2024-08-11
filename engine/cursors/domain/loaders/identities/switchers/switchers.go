package switchers

import (
	"errors"
	"fmt"
)

type switchersIns struct {
	list []Switcher
	mp   map[string]Switcher
}

func createSwitchers(
	list []Switcher,
	mp map[string]Switcher,
) Switchers {
	out := switchersIns{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list
func (obj *switchersIns) List() []Switcher {
	return obj.list
}

// Fetch fetches a switcher by name
func (obj *switchersIns) Fetch(name string) (Switcher, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the name (%s) could not be found", name)
	return nil, errors.New(str)
}
