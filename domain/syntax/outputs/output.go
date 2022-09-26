package outputs

import (
	"errors"
	"fmt"
)

type output struct {
	list []Variable
	mp   map[string]Variable
}

func createOutput(
	list []Variable,
	mp map[string]Variable,
) Output {
	out := output{
		list: list,
		mp:   mp,
	}

	return &out
}

// List returns the variables
func (obj *output) List() []Variable {
	return obj.list
}

// Find finds a variable by name
func (obj *output) Find(name string) (Variable, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the variable (name: %s) could not be found", name)
	return nil, errors.New(str)
}
