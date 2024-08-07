package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
)

type application struct {
	dbApp databases.Application
}

func createApplication(
	dbApp databases.Application,
) Application {
	out := application{
		dbApp: dbApp,
	}

	return &out
}

// Write writes the pointer
func (app *application) Write(pointer pointers.Pointer) error {
	storage := pointer.Storage()
	if storage.IsDeleted() {
		return nil
	}

	bytes := pointer.Bytes()
	index := storage.Delimiter().Index()
	return app.dbApp.CopyBeforeThenWrite(index, bytes)
}
