package inserts

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/inserts"
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
		code := failures.CouldNotFetchInstanceFromFrame
		return &code, err
	}

	pathVar := instruction.Path()
	path, err := frame.FetchStringList(pathVar)
	if err != nil {
		code := failures.CouldNotFetchPathFromFrame
		return &code, err
	}

	hash := instance.Hash()
	exists := app.repository.Exists(path, hash)
	if exists {
		code := failures.InstanceAlreadyExistsInDatabase
		pathStr := strings.Join(path, "/")
		str := fmt.Sprintf("the instance (path: %s, hash: %s) already exists in database", pathStr, hash.String())
		return &code, errors.New(str)
	}

	err = app.service.Insert(*pContext, instance, path)
	if err != nil {
		code := failures.CouldNotInsertInDatabase
		return &code, err
	}

	return nil, nil
}
