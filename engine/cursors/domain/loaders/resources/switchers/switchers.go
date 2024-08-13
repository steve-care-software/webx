package switchers

import (
	"errors"
	"fmt"
)

type switchers struct {
	list []Switcher
	mp   map[string]Switcher
}

func createSwitchers(
	list []Switcher,
	mp map[string]Switcher,
) Switchers {
	out := switchers{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of switchers
func (obj *switchers) List() []Switcher {
	return obj.list
}

// FetchByName fetches a switcher by name
func (obj *switchers) FetchByName(name string) (Switcher, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the switcher (name: %s) does not exists", name)
	return nil, errors.New(str)
}
