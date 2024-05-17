package databases

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	database_instruction "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	databaseService databases.Service
}

func createApplication(
	databaseService databases.Service,
) Application {
	out := application{
		databaseService: databaseService,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignment database_instruction.Database) (*uint, error) {
	if assignment.IsSave() {
		dbVar := assignment.Delete()
		db, err := frame.FetchDatabase(dbVar)
		if err != nil {
			code := failures.CouldNotFetchDatabaseFromFrame
			return &code, nil
		}

		err = app.databaseService.Save(db)
		if err != nil {
			code := failures.CouldNotSaveDatabaseFromService
			return &code, nil
		}

		return nil, err
	}

	delHashVar := assignment.Delete()
	hash, err := frame.FetchHash(delHashVar)
	if err != nil {
		code := failures.CouldNotFetchHashVariableFromFrame
		return &code, nil
	}

	err = app.databaseService.Delete(hash)
	if err != nil {
		code := failures.CouldNotDeleteDatabaseFromService
		return &code, nil
	}

	return nil, err
}
