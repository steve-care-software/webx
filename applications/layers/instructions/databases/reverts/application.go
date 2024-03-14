package reverts

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	service instances.Service
}

func createApplication(
	service instances.Service,
) Application {
	out := application{
		service: service,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, instruction reverts.Revert) (*uint, error) {
	if instruction.HasIndex() {
		indexVar := instruction.Index()
		pIndex, err := frame.FetchUnsignedInt(indexVar)
		if err != nil {
			code := failures.CouldNotFetchIndexFromFrame
			return &code, err
		}

		err = app.service.RevertToIndex(*pIndex)
		if err != nil {
			code := failures.CouldNotRevertToIndexInDatabase
			return &code, err
		}

		return nil, nil
	}

	err := app.service.Revert()
	if err != nil {
		code := failures.CouldNotRevertInDatabase
		return &code, err
	}

	return nil, nil
}
