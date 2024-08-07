package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	storage_pointers "github.com/steve-care-software/webx/engine/cursors/domain/storages/pointers"
)

type application struct {
	storagePointerBulder storage_pointers.PointerBuilder
	pointersBuilder      pointers.Builder
	pointerBuilder       pointers.PointerBuilder
	delimiterBuilder     delimiters.DelimiterBuilder
}

func createApplication(
	storagePointerBulder storage_pointers.PointerBuilder,
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

// InsertData inserts data to the pointers
func (app *application) InsertData(pointers pointers.Pointers, data []byte) (pointers.Pointers, error) {
	pNextIndex, err := pointers.NextIndex()
	if err != nil {
		return nil, err
	}

	length := uint64(len(data))
	delimiter, err := app.delimiterBuilder.Create().
		WithIndex(uint64(*pNextIndex)).
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
