package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers"
)

type storageHeaderAdapter struct {
	storagePointerAdapter storages.Adapter
	builder               headers.Builder
}

func createStorageHeaderAdapter(
	storagePointerAdapter storages.Adapter,
	builder headers.Builder,
) headers.Adapter {
	out := storageHeaderAdapter{
		storagePointerAdapter: storagePointerAdapter,
		builder:               builder,
	}

	return &out
}

// ToBytes converts header to bytes
func (app *storageHeaderAdapter) ToBytes(ins headers.Header) ([]byte, error) {
	output := []byte{
		HeaderFlag,
	}

	if ins.HasIdentities() {
		identities := ins.Identities()
		retBytes, err := app.storagePointerAdapter.InstanceToBytes(identities)
		if err != nil {
			return nil, err
		}

		lengthBytes := Uint64ToBytes(uint64(len(retBytes)))
		output = append(output, lengthBytes...)
		output = append(output, retBytes...)
	}

	return output, nil
}

// ToInstance converts bytes to instance
func (app *storageHeaderAdapter) ToInstance(data []byte) (headers.Header, error) {
	expectation := 1
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return nil, errors.New(str)
	}

	flag := data[0]
	if flag != HeaderFlag {
		return nil, errors.New("the data does not represents an Header instance, invalid flag")
	}

	builder := app.builder.Create()
	if len(data) > AmountOfBytesIntUint64 {
		toIndex := 1 + AmountOfBytesIntUint64
		length := int(BytesToUint64(data[1:toIndex]))
		identitiesBytes := data[toIndex : toIndex+length]
		identities, err := app.storagePointerAdapter.BytesToInstance(identitiesBytes)
		if err != nil {
			return nil, err
		}

		builder.WithIdentities(identities)
	}

	return builder.Now()
}
