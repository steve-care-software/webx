package mocks

import "github.com/steve-care-software/datastencil/domain/instances/databases"

type databaseRepository struct {
	list     [][]string
	database databases.Database
	errorIns error
}

func createDatabaseRepository(
	list [][]string,
	database databases.Database,
	errorIns error,
) databases.Repository {
	out := databaseRepository{
		list:     list,
		database: database,
		errorIns: errorIns,
	}

	return &out
}

// List returns the paths list
func (app *databaseRepository) List() ([][]string, error) {
	return app.list, app.errorIns
}

// Exists returns true if exists, false otherwise
func (app *databaseRepository) Exists(path []string) (*bool, error) {
	output := app.database != nil
	return &output, app.errorIns
}

// Retrieve retrieves a database from path
func (app *databaseRepository) Retrieve(path []string) (databases.Database, error) {
	return app.database, app.errorIns
}
