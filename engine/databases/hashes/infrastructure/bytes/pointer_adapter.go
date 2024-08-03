package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
	infra_bytes "github.com/steve-care-software/webx/engine/databases/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/pointers"
)

type pointerAdapter struct {
	hashAdapter      hash.Adapter
	delimiterAdapter delimiters.Adapter
	builder          pointers.Builder
	pointerBuilder   pointers.PointerBuilder
}

func createPointerAdapter(
	hashAdapter hash.Adapter,
	delimiterAdapter delimiters.Adapter,
	builder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
) pointers.Adapter {
	out := pointerAdapter{
		hashAdapter:      hashAdapter,
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

	lengthBytes := infra_bytes.Uint64ToBytes(uint64(len(list)))
	return append(lengthBytes, output...), nil
}

// BytesToInstances converts bytes to instances
func (app *pointerAdapter) BytesToInstances(data []byte) (pointers.Pointers, []byte, error) {
	amount, remaining, err := infra_bytes.AmountReturnRemaining(data)
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
	output := ins.Hash().Bytes()
	retBytes, err := app.delimiterAdapter.InstanceToBytes(ins.Delimiter())
	if err != nil {
		return nil, err
	}

	return append(output, retBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *pointerAdapter) BytesToInstance(data []byte) (pointers.Pointer, []byte, error) {
	if len(data) < hash.Size {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d bytes provided", hash.Size, len(data))
		return nil, nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(data[0:hash.Size])
	if err != nil {
		return nil, nil, err
	}

	retDelimiter, retRemaining, err := app.delimiterAdapter.BytesToInstance(data[hash.Size:])
	if err != nil {
		return nil, nil, err
	}

	pointer, err := app.pointerBuilder.Create().WithHash(*pHash).WithDelimiter(retDelimiter).Now()
	if err != nil {
		return nil, nil, err
	}

	return pointer, retRemaining, nil
}
