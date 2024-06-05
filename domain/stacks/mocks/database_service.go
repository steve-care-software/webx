package mocks

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
)

type databaseService struct {
	expected databases.Database
}

func createDatabaseService(
	expected databases.Database,
) databases.Service {
	out := databaseService{
		expected: expected,
	}

	return &out
}

// Save saves a database
func (app *databaseService) Save(database databases.Database) error {
	if app.expected.Hash().Compare(database.Hash()) {
		return nil
	}

	return errors.New("the database does not match the expectation")
}

// Delete deletes a database
func (app *databaseService) Delete(hash hash.Hash) error {
	if app.expected.Hash().Compare(hash) {
		return nil
	}

	return errors.New("the database hash does not exists")
}
