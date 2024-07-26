package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/databases/applications"
	"github.com/steve-care-software/webx/engine/databases/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/domain/headers/containers/pointers"
	"github.com/steve-care-software/webx/engine/databases/domain/states"
)

type application struct {
	stateAdapter   states.Adapter
	stateBuilder   states.Builder
	entriesBuilder entries.Builder
	deletesBuilder deletes.Builder
	contexts       map[uint]*context
}

func createApplication(
	stateAdapter states.Adapter,
	stateBuilder states.Builder,
	entriesBuilder entries.Builder,
	deletesBuilder deletes.Builder,
) applications.Application {
	out := application{
		stateAdapter:   stateAdapter,
		stateBuilder:   stateBuilder,
		entriesBuilder: entriesBuilder,
		deletesBuilder: deletesBuilder,
		contexts:       map[uint]*context{},
	}

	return &out
}

// Begin begins a context
func (app *application) Begin(path []string) (*uint, error) {
	filePath := filepath.Join(path...)
	pLock := fslock.New(filePath)
	err := pLock.TryLock()
	if err != nil {
		str := fmt.Sprintf("failed to acquire lock: %s", err.Error())
		return nil, errors.New(str)
	}

	pFile, err := os.Open(filePath)
	if err != nil {
		str := fmt.Sprintf("failed to open file: %s", err.Error())
		return nil, errors.New(str)
	}

	identifier := uint(len(app.contexts))
	app.contexts[identifier] = &context{
		path:       path,
		insertions: nil,
		deletions:  nil,
		pLock:      pLock,
		pFile:      pFile,
	}

	return &identifier, nil
}

// Retrieve retrieves entry data from a context
func (app *application) Retrieve(identifier uint, pointer pointers.Pointer) ([]byte, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		return app.readEntry(pContext.pFile, pointer)
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// RetrieveAll retrieves multiple entry data from context
func (app *application) RetrieveAll(identifier uint, pointers pointers.Pointers) ([][]byte, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		return app.readEntries(pContext.pFile, pointers)
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// Insert inserts an entry into a context
func (app *application) Insert(identifier uint, entry entries.Entry) error {
	if pContext, ok := app.contexts[identifier]; ok {
		entries, err := app.mergeInsert(pContext.insertions, []entries.Entry{
			entry,
		})

		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:       pContext.path,
			insertions: entries,
			deletions:  pContext.deletions,
			pLock:      pContext.pLock,
			pFile:      pContext.pFile,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// InsertAll inserts multiple entries into a context
func (app *application) InsertAll(identifier uint, newEntries entries.Entries) error {
	if pContext, ok := app.contexts[identifier]; ok {
		entries, err := app.mergeInsert(pContext.insertions, newEntries.List())

		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:       pContext.path,
			insertions: entries,
			deletions:  pContext.deletions,
			pLock:      pContext.pLock,
			pFile:      pContext.pFile,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Delete deletes an entry from a context
func (app *application) Delete(identifier uint, delete deletes.Delete) error {
	if pContext, ok := app.contexts[identifier]; ok {
		retDeletes, err := app.mergeDelete(pContext.deletions, []deletes.Delete{
			delete,
		})

		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:       pContext.path,
			insertions: pContext.insertions,
			deletions:  retDeletes,
			pLock:      pContext.pLock,
			pFile:      pContext.pFile,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// DeleteAll deletes multiple entries from from a context
func (app *application) DeleteAll(identifier uint, deletes deletes.Deletes) error {
	if pContext, ok := app.contexts[identifier]; ok {
		retDeletes, err := app.mergeDelete(pContext.deletions, deletes.List())
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:       pContext.path,
			insertions: pContext.insertions,
			deletions:  retDeletes,
			pLock:      pContext.pLock,
			pFile:      pContext.pFile,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Commit commits a context
func (app *application) Commit(context uint) error {
	return nil
}

// Rollback rollbacks a context
func (app *application) Rollback(context uint) error {
	return nil
}

// Cancel cancels a context
func (app *application) Cancel(identifier uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		err := pContext.pFile.Close()
		if err != nil {
			return err
		}

		err = pContext.pLock.Lock()
		if err != nil {
			return err
		}

		delete(app.contexts, identifier)
		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

func (app *application) readEntries(file *os.File, pointers pointers.Pointers) ([][]byte, error) {
	output := [][]byte{}
	list := pointers.List()
	for idx, onePointer := range list {
		bytes, err := app.readEntry(file, onePointer)
		if err != nil {
			str := fmt.Sprintf("could not read entry (pointer index: %d): %s", idx, err.Error())
			return nil, errors.New(str)
		}

		output = append(output, bytes)
	}

	return output, nil
}

func (app *application) readEntry(file *os.File, pointer pointers.Pointer) ([]byte, error) {
	index := pointer.Index()
	_, err := file.Seek(index, 0)
	if err != nil {
		return nil, err
	}

	length := pointer.Length()
	buffer := make([]byte, int64(length))
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (app *application) mergeInsert(original entries.Entries, newEntries []entries.Entry) (entries.Entries, error) {
	list := []entries.Entry{}
	if original != nil {
		list = append(list, original.List()...)
	}

	list = append(list, newEntries...)
	return app.entriesBuilder.Create().
		WithList(list).
		Now()
}

func (app *application) mergeDelete(original deletes.Deletes, newEntries []deletes.Delete) (deletes.Deletes, error) {
	list := []deletes.Delete{}
	if original != nil {
		list = append(list, original.List()...)
	}

	list = append(list, newEntries...)
	return app.deletesBuilder.Create().
		WithList(list).
		Now()
}
