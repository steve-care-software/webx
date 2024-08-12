package pointers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	storage_pointers "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type application struct {
	dbApp                databases.Application
	storagePointerBulder storage_pointers.StorageBuilder
	pointersBuilder      pointers.Builder
	pointerBuilder       pointers.PointerBuilder
	delimiterBuilder     delimiters.DelimiterBuilder
	nextIndex            uint64
}

func createApplication(
	dbApp databases.Application,
	storagePointerBulder storage_pointers.StorageBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	delimiterBuilder delimiters.DelimiterBuilder,
	nextIndex uint64,
) Application {
	out := application{
		dbApp:                dbApp,
		storagePointerBulder: storagePointerBulder,
		pointersBuilder:      pointersBuilder,
		pointerBuilder:       pointerBuilder,
		delimiterBuilder:     delimiterBuilder,
		nextIndex:            nextIndex,
	}

	return &out
}

// Retrieve retrieves a pointer from the database
func (app *application) Retrieve(storage storage_pointers.Storage) (pointers.Pointer, error) {
	return nil, nil
}

// InsertData inserts data to the pointers
func (app *application) InsertData(pointers pointers.Pointers, data []byte) (pointers.Pointers, error) {
	index := app.nextIndex
	length := uint64(len(data))
	delimiter, err := app.delimiterBuilder.Create().
		WithIndex(uint64(index)).
		WithLength(length).
		Now()

	if err != nil {
		return nil, err
	}

	storagePointer, err := app.storagePointerBulder.Create().
		WithDelimiter(delimiter).
		Now()

	if err != nil {
		return nil, err
	}

	pointer, err := app.pointerBuilder.Create().
		WithBytes(data).
		WithStorage(storagePointer).
		Now()

	if err != nil {
		return nil, err
	}

	list := pointers.List()
	list = append(list, pointer)
	return app.pointersBuilder.Create().
		WithList(list).
		Now()
}

// UpdateData updates data from the pointers
func (app *application) UpdateData(pointers pointers.Pointers, index uint64, updated []byte) (pointers.Pointers, error) {
	// delete the original:
	retPointers, err := app.DeleteData(pointers, index)
	if err != nil {
		return nil, err
	}

	// insert the updated
	return app.InsertData(retPointers, updated)
}

// DeleteData deletes data from the pointers
func (app *application) DeleteData(pointers pointers.Pointers, index uint64) (pointers.Pointers, error) {
	list := pointers.List()
	original := list[index]
	storage := original.Storage()
	if storage.IsDeleted() {
		str := fmt.Sprintf(pointerAlreadyDeletedErrPattern, index)
		return nil, errors.New(str)
	}

	delete := original.Storage().Delimiter()
	storagePointer, err := app.storagePointerBulder.Create().
		WithDelimiter(delete).
		IsDeleted().
		Now()

	if err != nil {
		return nil, err
	}

	data := original.Bytes()
	pointer, err := app.pointerBuilder.Create().
		WithBytes(data).
		WithStorage(storagePointer).
		Now()

	if err != nil {
		return nil, err
	}

	list[index] = pointer
	return app.pointersBuilder.Create().
		WithList(list).
		Now()
}

// Commit commits the data
func (app *application) Commit(pointers pointers.Pointers) error {
	return nil
}

// Purge purges the data
func (app *application) Purge(pointers pointers.Pointers) error {
	return nil
}

// PurgeAll purges all the data
func (app *application) PurgeAll(pointers pointers.Pointers) error {
	return nil
}

func (app *application) write(startAtIndex uint64, pointers pointers.Pointers) error {
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
