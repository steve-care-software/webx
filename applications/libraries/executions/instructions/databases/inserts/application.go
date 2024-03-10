package inserts

import (
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	repository instances.Repository
	service    instances.Service
}

func createApplication(
	repository instances.Repository,
	service instances.Service,
) Application {
	out := application{
		repository: repository,
		service:    service,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, instruction inserts.Insert) (*uint, error) {
	contextVar := instruction.Context()
	pContext, err := frame.FetchUnsignedInt(contextVar)
	if err != nil {
		code := failures.CouldNotFetchContextFromFrame
		return &code, err
	}

	insVar := instruction.Instance()
	instance, err := frame.FetchInstance(insVar)
	if err != nil {
		return nil, err
	}

	pathVar := instruction.Path()
	path, err := frame.FetchStringList(pathVar)
	if err != nil {
		code := failures.CouldNotFetchPathFromFrame
		return &code, err
	}

	hash := instance.Hash()
	exists, err := app.repository.Exists(path, hash)
	if err != nil {
		return nil, err
	}

	if exists {
		code := failures.InstanceAlreadyExists
		return &code, err
	}

	err = app.service.Insert(*pContext, instance, path)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
