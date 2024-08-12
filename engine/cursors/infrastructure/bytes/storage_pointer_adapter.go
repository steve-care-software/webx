package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
)

type storagePointerAdapter struct {
	delimiterAdapter delimiters.Adapter
	builder          storages.Builder
	storageBuilder   storages.StorageBuilder
}

func createStoragePointerAdapter(
	delimiterAdapter delimiters.Adapter,
	builder storages.Builder,
	storageBuilder storages.StorageBuilder,
) storages.Adapter {
	out := storagePointerAdapter{
		delimiterAdapter: delimiterAdapter,
		builder:          builder,
		storageBuilder:   storageBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *storagePointerAdapter) InstancesToBytes(ins storages.Storages) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, oneDelimiter := range list {
		retBytes, err := app.InstanceToBytes(oneDelimiter)
		if err != nil {
			return nil, err
		}

		lengthBytes := Uint64ToBytes(uint64(len(retBytes)))
		output = append(output, lengthBytes...)
		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instance
func (app *storagePointerAdapter) BytesToInstances(data []byte) (storages.Storages, error) {
	remaining := data
	list := []storages.Storage{}
	for {
		bytesLength, retRemaining, err := AmountReturnRemaining(remaining)
		if err != nil {
			return nil, err
		}

		ins, err := app.BytesToInstance(retRemaining[:bytesLength])
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
		remaining = retRemaining[bytesLength:]
		if len(remaining) <= 0 {
			break
		}
	}

	return app.builder.Create().WithList(list).Now()
}

// InstanceToBytes converts instance to bytes
func (app *storagePointerAdapter) InstanceToBytes(ins storages.Storage) ([]byte, error) {
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
func (app *storagePointerAdapter) BytesToInstance(data []byte) (storages.Storage, error) {
	expectation := 2
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return nil, errors.New(str)
	}

	flag := data[0]
	if flag != PointerFlag {
		return nil, errors.New("the data does not represents a Pointer instance, invalid flag")
	}

	isDeletedByte := data[1]
	builder := app.storageBuilder.Create()
	if isDeletedByte == 1 {
		builder.IsDeleted()
	}

	retDelimiter, err := app.delimiterAdapter.BytesToInstance(data[2:])
	if err != nil {
		return nil, err
	}

	return builder.WithDelimiter(retDelimiter).Now()
}
