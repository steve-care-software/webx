package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	storage_pointers "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

const pointerIsDeletedErrPattern = "the pointer (index: %d, length: %d) has been deleted"
const pointerAlreadyDeletedErrPattern = "the pointer (index: %d) has already been deleted"

// NewApplication creates a new application for tests
func NewApplication(
	dbApp databases.Application,
) Builder {
	storagePointerBulder := storage_pointers.NewStorageBuilder()
	pointersBuilder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	delimiterBuilder := delimiters.NewDelimiterBuilder()
	return createBuilder(
		dbApp,
		storagePointerBulder,
		pointersBuilder,
		pointerBuilder,
		delimiterBuilder,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithNextIndex(nextIndex uint64) Builder
	Now() (Application, error)
}

// Application represents the pointer application
type Application interface {
	Retrieve(storage storage_pointers.Storage) (pointers.Pointer, error)
	InsertData(pointers pointers.Pointers, data []byte) (pointers.Pointers, error)
	UpdateData(pointers pointers.Pointers, index uint64, updated []byte) (pointers.Pointers, error)
	DeleteData(pointers pointers.Pointers, index uint64) (pointers.Pointers, error)
	Commit(pointers pointers.Pointers) error
	Purge(pointers pointers.Pointers) error
	PurgeAll(pointers pointers.Pointers) error
}
