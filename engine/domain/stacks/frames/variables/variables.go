package variables

import (
	"errors"
	"fmt"
)

type variables struct {
	list []Variable
	mp   map[string]Variable
}

func createVariables(
	list []Variable,
	mp map[string]Variable,
) Variables {
	out := variables{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the list
func (obj *variables) List() []Variable {
	return obj.list
}

// Fetch fetches a variable by name
func (obj *variables) Fetch(name string) (Variable, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the variable (name: %s) does not exists", name)
	return nil, errors.New(str)
}
