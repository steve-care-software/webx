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

// Fetch fetches a pointer from the delimiter
func (obj *states) Fetch(delimiter delimiters.Delimiter) (pointers.Pointer, error) {
	length := len(obj.list)
	for i := 0; i < length; i++ {
		index := length - 1 - i
		if obj.list[index].IsDeleted() {
			continue
		}

		currentState := obj.list[index]
		if !currentState.HasPointers() {
			continue
		}

		retPointers, err := currentState.Pointers().Fetch(delimiter)
		if err != nil {
			continue
		}

		return retPointers, nil
	}

	return nil, errors.New("there is no pointers related to the provided delimiter")
}
