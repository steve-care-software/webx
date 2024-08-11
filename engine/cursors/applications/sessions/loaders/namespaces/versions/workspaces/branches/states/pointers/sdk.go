package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	storage_pointers "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

// NewApplication creates a new application for tests
func NewApplication() Application {
	storagePointerBulder := storage_pointers.NewStorageBuilder()
	pointersBuilder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	delimiterBuilder := delimiters.NewDelimiterBuilder()
	return createApplication(
		storagePointerBulder,
		pointersBuilder,
		pointerBuilder,
		delimiterBuilder,
	)
}

// Application represents the pointer application
type Application interface {
	InsertData(pointers pointers.Pointers, data []byte) (pointers.Pointers, error)
	UpdateData(pointers pointers.Pointers, original delimiters.Delimiter, updated []byte) (pointers.Pointers, error)
	DeleteData(pointers pointers.Pointers, delete delimiters.Delimiter) (pointers.Pointers, error)
}
