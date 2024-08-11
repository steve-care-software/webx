package loaders

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Application represents the loader application
type Application interface {
	Set(name string) (loaders.Loader, error)
	Down(name string) (loaders.Loader, error)
	Climb(name string) (loaders.Loader, error)
	Insert(original originals.Original) (loaders.Loader, error)
	Update(original string, updated originals.Original) (loaders.Loader, error)
	Delete(name string) (loaders.Loader, error)
	Recover(name string) (loaders.Loader, error)
	Purge(name string) (loaders.Loader, error)
	PurgeAll() (loaders.Loader, error)
	Move(name string, devName string, deleteOriginal bool) (loaders.Loader, error)
	Merge(deleteOriginal bool) (loaders.Loader, error)

	// data:
	InsertData(delimiter delimiters.Delimiter) (loaders.Loader, error)
	UpdateData(original delimiters.Delimiter, updated []byte) (loaders.Loader, error)
	DeleteData(delete delimiters.Delimiter) (loaders.Loader, error)
}
