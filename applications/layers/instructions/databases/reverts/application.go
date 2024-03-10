package reverts

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/reverts"
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

		}

		err = app.service.RevertToIndex(*pIndex)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	err := app.service.Revert()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
