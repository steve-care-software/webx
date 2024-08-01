package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

type delimiterAdapter struct {
	builder          delimiters.Builder
	delimiterBuilder delimiters.DelimiterBuilder
}

func createDelimiterAdapter(
	builder delimiters.Builder,
	delimiterBuilder delimiters.DelimiterBuilder,
) delimiters.Adapter {
	out := delimiterAdapter{
		builder:          builder,
		delimiterBuilder: delimiterBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *delimiterAdapter) InstancesToBytes(ins delimiters.Delimiters) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, oneDelimiter := range list {
		retBytes, err := app.InstanceToBytes(oneDelimiter)
		if err != nil {
			return nil, err
		}

		output = append(output, retBytes...)
	}

	lengthBytes := Uint64ToBytes(uint64(len(list)))
	return append(lengthBytes, output...), nil
}

// BytesToInstances converts bytes to instances
func (app *delimiterAdapter) BytesToInstances(data []byte) (delimiters.Delimiters, []byte, error) {
	amount, remaining, err := fetchAmountReturnRemaining(data)
	if err != nil {
		return nil, nil, err
	}

	input := remaining
	list := []delimiters.Delimiter{}
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
func (app *delimiterAdapter) InstanceToBytes(ins delimiters.Delimiter) ([]byte, error) {
	indexBytes := Uint64ToBytes(ins.Index())
	lengthBytes := Uint64ToBytes(ins.Length())
	return append(indexBytes, lengthBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *delimiterAdapter) BytesToInstance(data []byte) (delimiters.Delimiter, []byte, error) {
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
	ins, err := app.delimiterBuilder.Create().
		WithIndex(index).
		WithLength(length).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, data[toIndex:], nil
}
