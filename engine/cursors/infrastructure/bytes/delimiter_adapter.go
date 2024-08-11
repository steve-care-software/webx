package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
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

		lengthBytes := Uint64ToBytes(uint64(len(retBytes)))
		output = append(output, lengthBytes...)
		output = append(output, retBytes...)
	}

	return output, nil
}

// BytesToInstances converts bytes to instances
func (app *delimiterAdapter) BytesToInstances(data []byte) (delimiters.Delimiters, error) {
	remaining := data
	list := []delimiters.Delimiter{}
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
func (app *delimiterAdapter) InstanceToBytes(ins delimiters.Delimiter) ([]byte, error) {
	indexBytes := Uint64ToBytes(ins.Index())
	lengthBytes := Uint64ToBytes(ins.Length())
	output := append([]byte{
		DelimiterFlag,
	}, indexBytes...)
	return append(output, lengthBytes...), nil
}

// BytesToInstance converts bytes to instance
func (app *delimiterAdapter) BytesToInstance(data []byte) (delimiters.Delimiter, error) {
	expectation := (AmountOfBytesIntUint64 * 2) + 1
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return nil, errors.New(str)
	}

	flag := data[0]
	if flag != DelimiterFlag {
		return nil, errors.New("the data does not represents a Delimiter instance, invalid flag")
	}

	toIndex := AmountOfBytesIntUint64 + 1
	index := BytesToUint64(data[1:toIndex])

	fromIndex := toIndex
	toIndex = fromIndex + AmountOfBytesIntUint64
	length := BytesToUint64(data[fromIndex:])
	ins, err := app.delimiterBuilder.Create().
		WithIndex(index).
		WithLength(length).
		Now()

	if err != nil {
		return nil, err
	}

	return ins, nil
}
