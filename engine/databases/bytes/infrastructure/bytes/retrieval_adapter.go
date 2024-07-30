package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
)

type retrievalAdapter struct {
	builder          retrievals.Builder
	retrievalBuilder retrievals.RetrievalBuilder
}

func createRetrievalAdapter(
	builder retrievals.Builder,
	retrievalBuilder retrievals.RetrievalBuilder,
) retrievals.Adapter {
	out := retrievalAdapter{
		builder:          builder,
		retrievalBuilder: retrievalBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *retrievalAdapter) InstancesToBytes(ins retrievals.Retrievals) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, oneRetrieval := range list {
		retBytes, err := app.InstanceToBytes(oneRetrieval)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	lengthBytes := Uint64ToBytes(uint64(len(list)))
	return append(lengthBytes, output...), nil
}

// BytesToInstances converts bytes to instances
func (app *retrievalAdapter) BytesToInstances(data []byte) (retrievals.Retrievals, []byte, error) {
	amount, remaining, err := fetchAmountReturnRemaining(data)
	if err != nil {
		return nil, nil, err
	}

	input := remaining
	list := []retrievals.Retrieval{}
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
func (app *retrievalAdapter) InstanceToBytes(ins retrievals.Retrieval) ([]byte, error) {
	indexBytes := Uint64ToBytes(ins.Index())
	lengthBytes := Uint64ToBytes(ins.Length())
	return append(indexBytes, lengthBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *retrievalAdapter) BytesToInstance(data []byte) (retrievals.Retrieval, []byte, error) {
	expectation := AmountOfBytesIntUint64 * 2
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return nil, nil, errors.New(str)
	}

	toIndex := AmountOfBytesIntUint64
	index := BytesToUint64(data[:toIndex])

	fromIndex := toIndex
	toIndex = fromIndex + AmountOfBytesIntUint64
	length := BytesToUint64(data[fromIndex:])
	ins, err := app.retrievalBuilder.Create().
		WithIndex(index).
		WithLength(length).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, data[toIndex:], nil
}
