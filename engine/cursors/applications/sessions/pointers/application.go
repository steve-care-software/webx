package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	storage_pointers "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type application struct {
	storagePointerBulder storage_pointers.StorageBuilder
	pointersBuilder      pointers.Builder
	pointerBuilder       pointers.PointerBuilder
	delimiterBuilder     delimiters.DelimiterBuilder
}

func createApplication(
	storagePointerBulder storage_pointers.StorageBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	delimiterBuilder delimiters.DelimiterBuilder,
) Application {
	out := application{
		storagePointerBulder: storagePointerBulder,
		pointersBuilder:      pointersBuilder,
		pointerBuilder:       pointerBuilder,
		delimiterBuilder:     delimiterBuilder,
	}

	return &out
}

// Retrieve retrieves a pointer from the database
func (app *application) Retrieve(storage storage_pointers.Storage) (pointers.Pointer, error) {
	return nil, nil
}

// InsertData inserts data to the pointers
func (app *application) InsertData(pointers pointers.Pointers, data []byte) (pointers.Pointers, error) {
	index := pointers.NextIndex()
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
func (app *application) UpdateData(pointers pointers.Pointers, original delimiters.Delimiter, updated []byte) (pointers.Pointers, error) {
	// delete the original:
	retPointers, err := app.DeleteData(pointers, original)
	if err != nil {
		return nil, err
	}

	// insert the updated
	return app.InsertData(retPointers, updated)
}

// DeleteData deletes data from the pointers
func (app *application) DeleteData(pointers pointers.Pointers, delete delimiters.Delimiter) (pointers.Pointers, error) {
	pOriginalIndex, err := pointers.FindIndex(delete)
	if err != nil {
		return nil, err
	}

	originalIndex := *pOriginalIndex
	list := pointers.List()
	original := list[originalIndex]
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

	list[originalIndex] = pointer
	return app.pointersBuilder.Create().
		WithList(list).
		Now()
}
