package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
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
func (app *application) Write(startAtIndex uint64, pointers pointers.Pointers) error {
	cpyFromIndex := startAtIndex
	list := pointers.List()
	for _, onePointer := range list {
		storage := onePointer.Storage()
		if storage.IsDeleted() {
			continue
		}

		delimiter := storage.Delimiter()
		bytes := onePointer.Bytes()
		index := storage.Delimiter().Index()
		err := app.dbApp.CopyBeforeThenWrite(cpyFromIndex, index, bytes)
		if err != nil {
			return nil
		}

		cpyFromIndex = delimiter.Index() + delimiter.Length()
	}

	return nil
}
