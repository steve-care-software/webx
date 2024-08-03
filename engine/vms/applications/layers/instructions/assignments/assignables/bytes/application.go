package bytes

import (
	"bytes"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
	assignable_bytes "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

type application struct {
	hashAdapter       hash.Adapter
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	hashAdapter hash.Adapter,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		hashAdapter:       hashAdapter,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable assignable_bytes.Bytes) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsJoin() {
		output := []byte{}
		variables := assignable.Join()
		for _, oneVariable := range variables {
			data, err := frame.FetchBytes(oneVariable)
			if err != nil {
				code := failures.CouldNotFetchJoinVariableFromFrame
				return nil, &code, err
			}

			output = append(output, data...)
		}

		builder.WithBytes(output)
	}

	if assignable.IsCompare() {
		boolValue := true
		var firstBytes []byte
		variables := assignable.Compare()
		for _, oneVariable := range variables {
			data, err := frame.FetchBytes(oneVariable)
			if err != nil {
				code := failures.CouldNotFetchCompareVariableFromFrame
				return nil, &code, err
			}

			if firstBytes == nil {
				firstBytes = data
				continue
			}

			if !bytes.Equal(firstBytes, data) {
				boolValue = false
				break
			}
		}

		builder.WithBool(boolValue)
	}

	if assignable.IsHashBytes() {
		variable := assignable.HashBytes()
		data, err := frame.FetchBytes(variable)
		if err != nil {
			code := failures.CouldNotFetchHashVariableFromFrame
			return nil, &code, err
		}

		pHash, err := app.hashAdapter.FromBytes(data)
		if err != nil {
			return nil, nil, err
		}

		builder.WithHash(*pHash)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
