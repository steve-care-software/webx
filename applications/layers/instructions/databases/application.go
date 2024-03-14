package databases

import (
	application_deletes "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/deletes"
	application_inserts "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/inserts"
	application_reverts "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execDeleteApp application_deletes.Application
	execInsertApp application_inserts.Application
	execRevertApp application_reverts.Application
	service       instances.Service
}

func createApplication(
	execDeleteApp application_deletes.Application,
	execInsertApp application_inserts.Application,
	execRevertApp application_reverts.Application,
	service instances.Service,
) Application {
	out := application{
		execDeleteApp: execDeleteApp,
		execInsertApp: execInsertApp,
		execRevertApp: execRevertApp,
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
			code := failures.CouldNotCommitInDatabase
			return &code, err
		}

		return nil, nil
	}

	cancelVar := instruction.Cancel()
	pContext, err := frame.FetchUnsignedInt(cancelVar)
	if err != nil {
		code := failures.CouldNotFetchContextFromFrame
		return &code, err
	}

	err = app.service.Cancel(*pContext)
	if err != nil {
		code := failures.CouldNotCancelInDatabase
		return &code, err
	}

	return nil, nil
}
