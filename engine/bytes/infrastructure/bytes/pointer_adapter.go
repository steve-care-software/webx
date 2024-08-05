package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/pointers"
)

type pointerAdapter struct {
	delimiterAdapter delimiters.Adapter
	builder          pointers.Builder
	pointerBuilder   pointers.PointerBuilder
}

func createPointerAdapter(
	delimiterAdapter delimiters.Adapter,
	builder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
) pointers.Adapter {
	out := pointerAdapter{
		delimiterAdapter: delimiterAdapter,
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
	amount, remaining, err := AmountReturnRemaining(data)
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
	retrievalBytes, err := app.delimiterAdapter.InstanceToBytes(ins.Delimiter())
	if err != nil {
		return nil, err
	}

	output := []byte{
		PointerFlag,
		0, // 1 == true, 0 == false
	}
	if ins.IsDeleted() {
		output[1] = 1
	}

	return append(output, retrievalBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *pointerAdapter) BytesToInstance(data []byte) (pointers.Pointer, []byte, error) {
	expectation := 2
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return nil, nil, errors.New(str)
	}

	flag := data[0]
	if flag != PointerFlag {
		return nil, nil, errors.New("the data does not represents a Pointer instance, invalid flag")
	}

	isDeletedByte := data[1]
	builder := app.pointerBuilder.Create()
	if isDeletedByte == 1 {
		builder.IsDeleted()
	}

	retDelimiter, retRemaining, err := app.delimiterAdapter.BytesToInstance(data[2:])
	if err != nil {
		return nil, nil, err
	}

	ins, err := builder.WithDelimiter(retDelimiter).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, retRemaining, nil
}
