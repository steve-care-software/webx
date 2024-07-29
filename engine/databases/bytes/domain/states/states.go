package states

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
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

// Amount returns the amount of pointers a containers contains
func (obj *states) Amount(keyname string) (*uint, error) {
	container, err := obj.fetchContainer(keyname)
	if err != nil {
		return nil, err
	}

	list := container.Pointers().List()
	amount := uint(len(list))
	return &amount, nil
}

// Fetch fetches a list of retrievals based on a container
func (obj *states) Fetch(keyname string, index uint64, length uint64) ([]retrievals.Retrieval, error) {
	container, err := obj.fetchContainer(keyname)
	if err != nil {
		return nil, err
	}

	return container.Pointers().Fetch(index, length)
}

func (obj *states) fetchContainer(keyname string) (containers.Container, error) {
	currentState, err := obj.currentState()
	if err != nil {
		return nil, err
	}

	if !currentState.HasContainers() {
		return nil, errors.New("the current state contains no contaniners")
	}

	return currentState.Containers().Fetch(keyname)
}

func (obj *states) currentState() (State, error) {
	length := len(obj.list)
	for i := 0; i < length; i++ {
		index := length - i
		if obj.list[index].IsDeleted() {
			continue
		}

		return obj.list[index], nil
	}

	return nil, errors.New("there is no current active state")
}
