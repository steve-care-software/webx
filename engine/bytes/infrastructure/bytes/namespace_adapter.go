package bytes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
)

type namespaceAdapter struct {
	delimiterAdapter delimiters.Adapter
	builder          namespaces.Builder
	namespaceBuilder namespaces.NamespaceBuilder
}

func createNamespaceAdapter(
	delimiterAdapter delimiters.Adapter,
	builder namespaces.Builder,
	namespaceBuilder namespaces.NamespaceBuilder,
) namespaces.Adapter {
	out := namespaceAdapter{
		delimiterAdapter: delimiterAdapter,
		builder:          builder,
		namespaceBuilder: namespaceBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *namespaceAdapter) InstancesToBytes(ins namespaces.Namespaces) ([]byte, error) {
	output := []byte{}
	list := ins.List()
	for _, oneNamespace := range list {
		retBytes, err := app.InstanceToBytes(oneNamespace)
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
func (app *namespaceAdapter) BytesToInstances(data []byte) (namespaces.Namespaces, error) {
	remaining := data
	list := []namespaces.Namespace{}
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
func (app *namespaceAdapter) InstanceToBytes(ins namespaces.Namespace) ([]byte, error) {
	nameBytes := []byte(ins.Name())
	nameLengthBytes := Uint64ToBytes(uint64(len(nameBytes)))

	descriptionBytes := []byte(ins.Description())
	descriptionLengthBytes := Uint64ToBytes(uint64(len(descriptionBytes)))

	isDeletedBytes := []byte{
		0,
	}

	if ins.IsDeleted() {
		isDeletedBytes[0] = 1
	}

	output := []byte{
		NamespaceFlag,
	}

	output = append(output, nameLengthBytes...)
	output = append(output, nameBytes...)
	output = append(output, descriptionLengthBytes...)
	output = append(output, descriptionBytes...)
	output = append(output, isDeletedBytes...)
	if ins.HasIterations() {
		delimiterBytes, err := app.delimiterAdapter.InstanceToBytes(ins.Iterations())
		if err != nil {
			return nil, err
		}

		delimiterLengthBytes := Uint64ToBytes(uint64(len(delimiterBytes)))

		output = append(output, delimiterLengthBytes...)
		output = append(output, delimiterBytes...)
	}

	return output, nil
}

// BytesToInstance converts bytes to instance
func (app *namespaceAdapter) BytesToInstance(data []byte) (namespaces.Namespace, error) {
	// flag:
	remaining := data
	expectation := 1
	if len(remaining) < expectation {
		str := fmt.Sprintf(
			"flag: %s",
			fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
		)

		return nil, errors.New(str)
	}

	flag := remaining[0]
	if flag != NamespaceFlag {
		return nil, errors.New("the data does not represents a Namespace instance, invalid flag")
	}

	remaining = remaining[expectation:]

	// name length:
	expectation = AmountOfBytesIntUint64
	if len(remaining) < expectation {
		str := fmt.Sprintf(
			"name length: %s",
			fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
		)

		return nil, errors.New(str)
	}

	nameLength := BytesToUint64(remaining[:expectation])
	remaining = remaining[expectation:]

	// name:
	expectation = int(nameLength)
	if len(remaining) < expectation {
		str := fmt.Sprintf(
			"name: %s",
			fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
		)

		return nil, errors.New(str)
	}

	name := string(remaining[:expectation])
	remaining = remaining[expectation:]
	builder := app.namespaceBuilder.Create().WithName(name)

	// description length:
	expectation = AmountOfBytesIntUint64
	if len(remaining) < expectation {
		str := fmt.Sprintf(
			"description length: %s",
			fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
		)

		return nil, errors.New(str)
	}

	descriptionLength := BytesToUint64(remaining[:expectation])
	remaining = remaining[expectation:]

	// description:
	expectation = int(descriptionLength)
	if len(remaining) < expectation {
		str := fmt.Sprintf(
			"description: %s",
			fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
		)

		return nil, errors.New(str)
	}

	description := string(remaining[:expectation])
	remaining = remaining[expectation:]
	builder.WithDescription(description)

	// isDeleted:
	expectation = 1
	if len(remaining) < expectation {
		str := fmt.Sprintf(
			"isDeleted: %s",
			fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
		)

		return nil, errors.New(str)
	}

	// true
	if remaining[0] == 1 {
		builder.IsDeleted()
	}

	remaining = remaining[expectation:]

	// iterations:
	if len(remaining) > 0 {
		// iteration length:
		expectation = AmountOfBytesIntUint64
		if len(remaining) < expectation {
			str := fmt.Sprintf(
				"iteration length: %s",
				fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
			)

			return nil, errors.New(str)
		}

		iterationLength := BytesToUint64(remaining[:expectation])
		remaining = remaining[expectation:]

		// iteration:
		expectation = int(iterationLength)
		if len(remaining) < expectation {
			str := fmt.Sprintf(
				"iteration: %s",
				fmt.Sprintf(invalidDataPatternErr, expectation, len(remaining)),
			)

			return nil, errors.New(str)
		}

		iteration, _, err := app.delimiterAdapter.BytesToInstance(remaining[:expectation])
		if err != nil {
			return nil, err
		}

		builder.WithIterations(iteration)
	}

	return builder.Now()
}
