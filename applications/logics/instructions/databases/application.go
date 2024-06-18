package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	database_instruction "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
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
		dbVar := assignment.Save()
		db, err := frame.FetchDatabase(dbVar)
		if err != nil {
			code := failures.CouldNotFetchDatabaseFromFrame
			return &code, err
		}

		err = app.databaseService.Save(db)
		if err != nil {
			code := failures.CouldNotSaveDatabaseFromService
			return &code, err
		}

		return nil, err
	}

	delHashVar := assignment.Delete()
	hash, err := frame.FetchHash(delHashVar)
	if err != nil {
		code := failures.CouldNotFetchHashVariableFromFrame
		return &code, err
	}

	err = app.databaseService.Delete(hash)
	if err != nil {
		code := failures.CouldNotDeleteDatabaseFromService
		return &code, err
	}

	return nil, err
}
