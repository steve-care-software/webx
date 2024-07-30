package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
)

type pointerAdapter struct {
	retrievalAdapter retrievals.Adapter
	builder          pointers.Builder
	pointerBuilder   pointers.PointerBuilder
}

func createPointerAdapter(
	retrievalAdapter retrievals.Adapter,
	builder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
) pointers.Adapter {
	out := pointerAdapter{
		retrievalAdapter: retrievalAdapter,
		builder:          builder,
		pointerBuilder:   pointerBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *pointerAdapter) InstancesToBytes(ins pointers.Pointers) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, onePointer := range list {
		retBytes, err := app.InstanceToBytes(onePointer)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	lengthBytes := Uint64ToBytes(uint64(len(list)))
	return append(lengthBytes, output...), nil
}

// BytesToInstances converts bytes to instances
func (app *pointerAdapter) BytesToInstances(data []byte) (pointers.Pointers, []byte, error) {
	amount, remaining, err := fetchAmountReturnRemaining(data)
	if err != nil {
		return nil, nil, err
	}

	input := remaining
	list := []pointers.Pointer{}
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
func (app *pointerAdapter) InstanceToBytes(ins pointers.Pointer) ([]byte, error) {
	retrievalBytes, err := app.retrievalAdapter.InstanceToBytes(ins.Retrieval())
	if err != nil {
		return nil, err
	}

	// 1 == true, 0 == false
	output := []byte{0}
	if ins.IsDeleted() {
		output[0] = 1
	}

	return append(output, retrievalBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *pointerAdapter) BytesToInstance(data []byte) (pointers.Pointer, []byte, error) {
	expectation := 1
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return nil, nil, errors.New(str)
	}

	isDeletedByte := data[0]
	builder := app.pointerBuilder.Create()
	if isDeletedByte == 1 {
		builder.IsDeleted()
	}

	retRetrieval, retRemaining, err := app.retrievalAdapter.BytesToInstance(data[1:])
	if err != nil {
		return nil, nil, err
	}

	ins, err := builder.WithRetrieval(retRetrieval).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}
