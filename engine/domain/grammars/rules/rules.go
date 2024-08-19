package rules

import (
	"errors"
	"fmt"
)

type rules struct {
	list []Rule
	mp   map[string]Rule
}

func createRules(
	list []Rule,
	mp map[string]Rule,
) Rules {
	out := rules{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list of rule
func (obj *rules) List() []Rule {
	return obj.list
}

// Fetch fetches a rule by name
func (obj *rules) Fetch(name string) (Rule, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the rule (name: %s) does not exists", name)
	return nil, errors.New(str)
}
