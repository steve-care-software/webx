package entries

import (
	"github.com/steve-care-software/webx/domain/databases"
)

type application struct {
	database databases.Database
}

func createAppliction(
	database databases.Database,
) Application {
	out := application{
		database: database,
	}

	return &out
}

// Retrieve retrieves an entry by pointer
func (app *application) Retrieve(pointer databases.Pointer) (databases.Entry, error) {
	return nil, nil
}

// Insert inserts an entry and returns its pointer
func (app *application) Insert(entry databases.Entry) (databases.Pointer, error) {
	// retrieve the section for identities:

	// retrieve the next pointer index:

	// save the transaction, if any:

	// save the entry instance:
	return nil, nil
}
