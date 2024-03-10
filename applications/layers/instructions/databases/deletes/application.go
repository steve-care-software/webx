package deletes

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases/deletes"
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
func (app *application) Execute(frame stacks.Frame, instruction deletes.Delete) (*uint, error) {
	contextVar := instruction.Context()
	pContext, err := frame.FetchUnsignedInt(contextVar)
	if err != nil {
		code := failures.CouldNotFetchContextFromFrame
		return &code, err
	}

	pathVar := instruction.Path()
	path, err := frame.FetchStringList(pathVar)
	if err != nil {
		code := failures.CouldNotFetchPathFromFrame
		return &code, err
	}

	hashVar := instruction.Identifier()
	hash, err := frame.FetchHash(hashVar)
	if err != nil {
		code := failures.CouldNotFetchIdentifierFromFrame
		return &code, err
	}

	exists, err := app.repository.Exists(path, hash)
	if err != nil {
		return nil, err
	}

	if !exists {
		code := failures.InstanceDoesNotExists
		return &code, err
	}

	err = app.service.Delete(*pContext, path, hash)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
