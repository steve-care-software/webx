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
	"github.com/steve-care-software/webx/engine/databases/domain/headers"
	"github.com/steve-care-software/webx/engine/databases/domain/modifications"
	"github.com/steve-care-software/webx/engine/databases/domain/retrievals"
)

type application struct {
	headerAdapter       headers.Adaptetr
	modificationAdapter modifications.Adapter
	modificationBuilder modifications.Builder
	entriesBuilder      entries.Builder
	deletesBuilder      deletes.Builder
	retrievalsBuilder   retrievals.Builder
	contexts            map[uint]*context
}

func createApplication(
	headerAdapter headers.Adaptetr,
	modificationAdapter modifications.Adapter,
	modificationBuilder modifications.Builder,
	entriesBuilder entries.Builder,
	deletesBuilder deletes.Builder,
	retrievalsBuilder retrievals.Builder,
) applications.Application {
	out := application{
		headerAdapter:       headerAdapter,
		modificationAdapter: modificationAdapter,
		modificationBuilder: modificationBuilder,
		entriesBuilder:      entriesBuilder,
		deletesBuilder:      deletesBuilder,
		retrievalsBuilder:   retrievalsBuilder,
		contexts:            map[uint]*context{},
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

	currentHeader, err := app.readHeader(pFile)
	if err != nil {
		str := fmt.Sprintf("failed to read Header: %s", err.Error())
		return nil, errors.New(str)
	}

	identifier := uint(len(app.contexts))
	app.contexts[identifier] = &context{
		path:          path,
		currentHeader: currentHeader,
		insertions:    nil,
		deletions:     nil,
		pLock:         pLock,
		pFile:         pFile,
	}

	return &identifier, nil
}

// List returns the list of pointers
func (app *application) List(identifier uint, keyname string, index uint, length uint) (retrievals.Retrievals, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		list, err := pContext.currentHeader.Fetch(keyname, index, length)
		if err != nil {
			return nil, err
		}

		return app.retrievalsBuilder.Create().
			WithList(list).
			Now()
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// Amount returns the amount of entities in the keyanme
func (app *application) Amount(identifier uint, keyname string) (*uint, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		return pContext.currentHeader.Amount(keyname)
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// Retrieve retrieves entry data from a context
func (app *application) Retrieve(identifier uint, retrieval retrievals.Retrieval) ([]byte, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		return app.readEntry(pContext.pFile, retrieval)
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// RetrieveAll retrieves multiple entry data from context
func (app *application) RetrieveAll(identifier uint, retrievals retrievals.Retrievals) ([][]byte, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		return app.readEntries(pContext.pFile, retrievals)
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

// Rollback rollbacks a context to the previous state
func (app *application) Rollback(context uint) error {
	return nil
}

// RollbackTo rollbacks a context to the amount provided
func (app *application) RollbackTo(context uint, amount uint) error {
	return nil
}

// RollFront rollfronts a context to the front state state
func (app *application) RollFront(context uint) error {
	return nil
}

// RollFrontTo rollfronts a context to the amount provided
func (app *application) RollFrontTo(context uint, amount uint) error {
	return nil
}

// States returns the amount of states
func (app *application) States(context uint, includesDeleted bool) (*uint, error) {
	return nil, nil
}

// DeletedStates returns the amount of deleted states
func (app *application) DeletedStates(context uint) (*uint, error) {
	return nil, nil
}

// Purge purges the previous states and only keep the latest one.  It also deletes previously deleted entries
func (app *application) Purge(context uint) error {
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

func (app *application) readEntries(file *os.File, retrievals retrievals.Retrievals) ([][]byte, error) {
	output := [][]byte{}
	list := retrievals.List()
	for idx, oneRetrieval := range list {
		bytes, err := app.readEntry(file, oneRetrieval)
		if err != nil {
			str := fmt.Sprintf("could not read entry (pointer index: %d): %s", idx, err.Error())
			return nil, errors.New(str)
		}

		output = append(output, bytes)
	}

	return output, nil
}

func (app *application) readHeader(file *os.File) (headers.Header, error) {
	// read the first int64 of the file:
	lengthBytes, err := app.readBytes(file, 0, amountOfBytesIntUint64)
	if err != nil {
		return nil, err
	}

	// convert the bytes to the length:
	length := int64(bytesToUint64(lengthBytes))

	// read the data:
	headerBytes, err := app.readBytes(file, amountOfBytesIntUint64, length)
	if err != nil {
		return nil, err
	}

	return app.headerAdapter.ToInstance(headerBytes)
}

func (app *application) readEntry(file *os.File, retrieval retrievals.Retrieval) ([]byte, error) {
	index := retrieval.Index()
	length := retrieval.Length()
	return app.readBytes(file, index, length)
}

func (app *application) readBytes(file *os.File, index int64, length int64) ([]byte, error) {
	_, err := file.Seek(index, 0)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, length)
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
