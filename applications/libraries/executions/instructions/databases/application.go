package databases

import (
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execDeleteApp deletes.Application
	execInsertApp inserts.Application
	execRevertApp reverts.Application
	repository    instances.Repository
	service       instances.Service
}

func createApplication(
	execDeleteApp deletes.Application,
	execInsertApp inserts.Application,
	execRevertApp reverts.Application,
	repository instances.Repository,
	service instances.Service,
) Application {
	out := application{
		execDeleteApp: execDeleteApp,
		execInsertApp: execInsertApp,
		execRevertApp: execRevertApp,
		repository:    repository,
		service:       service,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, instruction databases.Database) (*uint, error) {
	if instruction.IsInsert() {
		insert := instruction.Insert()
		return app.execInsertApp.Execute(frame, insert)
	}

	if instruction.IsDelete() {
		delete := instruction.Delete()
		return app.execDeleteApp.Execute(frame, delete)
	}

	if instruction.IsRevert() {
		revert := instruction.Revert()
		return app.execRevertApp.Execute(frame, revert)
	}

	if instruction.IsCommit() {
		commitVar := instruction.Commit()
		pContext, err := frame.FetchUnsignedInt(commitVar)
		if err != nil {
			code := failures.CouldNotFetchContextFromFrame
			return &code, err
		}

		err = app.service.Commit(*pContext)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	commitVar := instruction.Cancel()
	pContext, err := frame.FetchUnsignedInt(commitVar)
	if err != nil {
		code := failures.CouldNotFetchContextFromFrame
		return &code, err
	}

	err = app.service.Cancel(*pContext)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
