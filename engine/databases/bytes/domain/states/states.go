package states

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

type states struct {
	list []State
}

func createStates(
	list []State,
) States {
	out := states{
		list: list,
	}

	return &out
}

// List returns the list of states
func (obj *states) List() []State {
	return obj.list
}

// Fetch fetches a pointer from the keyname and delimiter
func (obj *states) Fetch(delimiter delimiters.Delimiter) (pointers.Pointer, error) {
	pointers, err := obj.fetchCurrentStatePointers()
	if err != nil {
		return nil, err
	}

	return pointers.Fetch(delimiter)
}

// Subset fetches a list subset of pointers based on a container
func (obj *states) Subset(index uint64, length uint64) ([]pointers.Pointer, error) {
	pointers, err := obj.fetchCurrentStatePointers()
	if err != nil {
		return nil, err
	}

	return pointers.Subset(index, length)
}

func (obj *states) fetchCurrentStatePointers() (pointers.Pointers, error) {
	currentState, err := obj.currentState()
	if err != nil {
		return nil, err
	}

	if !currentState.HasPointers() {
		return nil, errors.New("the current state contains no pointers")
	}

	return currentState.Pointers(), nil
}

func (obj *states) currentState() (State, error) {
	length := len(obj.list)
	for i := 0; i < length; i++ {
		index := length - 1 - i
		if obj.list[index].IsDeleted() {
			continue
		}

		return obj.list[index], nil
	}

	return nil, errors.New("there is no current active state")
}
