package stacks

import (
	"errors"
	"fmt"
)

type assignments struct {
	list []Assignment
}

func createAssignments(
	list []Assignment,
) Assignments {
	out := assignments{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *assignments) List() []Assignment {
	return obj.list
}

// Fetch fetches an assignable by name
func (obj *assignments) Fetch(name string) (Assignable, error) {
	for _, oneAssignment := range obj.list {
		if oneAssignment.Name() != name {
			continue
		}

		return oneAssignment.Assignable(), nil
	}

	str := fmt.Sprintf("there is no assignable related to the provided name: %s", name)
	return nil, errors.New(str)
}
