package bytes

import (
	"bytes"

	"github.com/steve-care-software/datastencil/domain/hash"
	assignable_bytes "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
func (app *application) Execute(frame stacks.Frame, assignable assignable_bytes.Bytes) (stacks.Assignable, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsJoin() {
		output := []byte{}
		variables := assignable.Join()
		for _, oneVariable := range variables {
			data, err := frame.FetchBytes(oneVariable)
			if err != nil {
				return nil, err
			}

			output = append(output, data...)
		}

		builder.WithBytes(output)
	}

	if assignable.IsCompare() {
		boolValue := true
		var lastBytes []byte
		variables := assignable.Compare()
		for _, oneVariable := range variables {
			data, err := frame.FetchBytes(oneVariable)
			if err != nil {
				return nil, err
			}

			if lastBytes == nil {
				lastBytes = data
				continue
			}

			if !bytes.Equal(lastBytes, data) {
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
			return nil, err
		}

		pHash, err := app.hashAdapter.FromBytes(data)
		if err != nil {
			return nil, err
		}

		builder.WithHash(*pHash)
	}

	return builder.Now()
}
