package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
)

type stateAdapter struct {
	containerAdapter containers.Adapter
	builder          states.Builder
	stateBuilder     states.StateBuilder
}

func createStateAdapter(
	containerAdapter containers.Adapter,
	builder states.Builder,
	stateBuilder states.StateBuilder,
) states.Adapter {
	out := stateAdapter{
		containerAdapter: containerAdapter,
		builder:          builder,
		stateBuilder:     stateBuilder,
	}

	return &out
}

// InstancesToBytes converts an instances to bytes
func (app *stateAdapter) InstancesToBytes(ins states.States) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, oneState := range list {
		retBytes, err := app.InstanceToBytes(oneState)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	lengthBytes := Uint64ToBytes(uint64(len(list)))
	return append(lengthBytes, output...), nil
}

// BytesToInstances converts bytes to instances
func (app *stateAdapter) BytesToInstances(data []byte) (states.States, []byte, error) {
	amount, remaining, err := fetchAmountReturnRemaining(data)
	if err != nil {
		return nil, nil, err
	}

	input := remaining
	list := []states.State{}
	for i := 0; i < amount; i++ {
		retIns, retRemaining, err := app.BytesToInstance(input)
		if err != nil {
			break
		}

		input = retRemaining
		list = append(list, retIns)
	}

	ins, err := app.builder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, input, nil
}

// InstanceToBytes converts an instance to bytes
func (app *stateAdapter) InstanceToBytes(ins states.State) ([]byte, error) {
	// 1 == true, 0 == false
	output := []byte{0}
	if ins.IsDeleted() {
		output[0] = 1
	}

	if ins.HasContainers() {
		retBytes, err := app.containerAdapter.InstancesToBytes(ins.Containers())
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstance converts bytes to instance
func (app *stateAdapter) BytesToInstance(data []byte) (states.State, []byte, error) {
	expectation := 1
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return nil, nil, errors.New(str)
	}

	isDeletedByte := data[0]
	builder := app.stateBuilder.Create()
	if isDeletedByte == 1 {
		builder.IsDeleted()
	}

	containers, retRemaining, err := app.containerAdapter.BytesToInstances(data[1:])
	if err != nil {
		retRemaining = data[1:]
	}

	if containers != nil {
		builder.WithContainers(containers)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}
