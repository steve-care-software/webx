package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
)

type containerAdapter struct {
	pointerAdapter   pointers.Adapter
	builder          containers.Builder
	containerBuilder containers.ContainerBuilder
}

func createContainerAdapter(
	pointerAdapter pointers.Adapter,
	builder containers.Builder,
	containerBuilder containers.ContainerBuilder,
) containers.Adapter {
	out := containerAdapter{
		pointerAdapter:   pointerAdapter,
		builder:          builder,
		containerBuilder: containerBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *containerAdapter) InstancesToBytes(ins containers.Containers) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, oneContainer := range list {
		retBytes, err := app.InstanceToBytes(oneContainer)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	lengthBytes := Uint64ToBytes(uint64(len(list)))
	return append(lengthBytes, output...), nil
}

// BytesToInstances converts bytes to instances
func (app *containerAdapter) BytesToInstances(data []byte) (containers.Containers, []byte, error) {
	amount, remaining, err := fetchAmountReturnRemaining(data)
	if err != nil {
		return nil, nil, err
	}

	input := remaining
	list := []containers.Container{}
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

// InstanceToBytes converts instance to bytes
func (app *containerAdapter) InstanceToBytes(ins containers.Container) ([]byte, error) {
	keynameBytes := []byte(ins.Keyname())
	keynameLengthBytes := Uint64ToBytes(uint64(len(keynameBytes)))
	pointersBytes, err := app.pointerAdapter.InstancesToBytes(ins.Pointers())
	if err != nil {
		return nil, err
	}

	output := append(keynameLengthBytes, keynameBytes...)
	return append(output, pointersBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *containerAdapter) BytesToInstance(data []byte) (containers.Container, []byte, error) {
	keynameAmount, remaining, err := fetchAmountReturnRemaining(data)
	if err != nil {
		return nil, nil, err
	}

	if len(remaining) < keynameAmount {
		str := fmt.Sprintf("%d bytes were expected in the data to extract the keyname, data length provided contains %d bytes", keynameAmount, len(remaining))
		return nil, nil, errors.New(str)
	}

	keyname := string(remaining[:keynameAmount])
	pointers, retRemaining, err := app.pointerAdapter.BytesToInstances(remaining[keynameAmount:])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.containerBuilder.Create().
		WithKeyname(keyname).
		WithPointers(pointers).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}
