package files

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/databases/bytes/applications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
	infra_bytes "github.com/steve-care-software/webx/engine/databases/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type application struct {
	hashAdapter       hash.Adapter
	statesAdapter     states.Adapter
	statesBuilder     states.Builder
	stateBuilder      states.StateBuilder
	pointersBuilder   pointers.Builder
	pointerBuilder    pointers.PointerBuilder
	entriesBuilder    entries.Builder
	delimitersBuilder delimiters.Builder
	delimiterBuilder  delimiters.DelimiterBuilder
	basepath          []string
	contexts          map[uint]*context
}

func createApplication(
	hashAdapter hash.Adapter,
	statesAdapter states.Adapter,
	statesBuilder states.Builder,
	stateBuilder states.StateBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	entriesBuilder entries.Builder,
	delimitersBuilder delimiters.Builder,
	delimiterBuilder delimiters.DelimiterBuilder,
	basepath []string,
) applications.Application {
	out := application{
		hashAdapter:       hashAdapter,
		statesAdapter:     statesAdapter,
		statesBuilder:     statesBuilder,
		stateBuilder:      stateBuilder,
		pointersBuilder:   pointersBuilder,
		pointerBuilder:    pointerBuilder,
		entriesBuilder:    entriesBuilder,
		delimitersBuilder: delimitersBuilder,
		delimiterBuilder:  delimiterBuilder,
		basepath:          basepath,
		contexts:          map[uint]*context{},
	}

	return &out
}

// Begin begins a context
func (app *application) Begin(name string) (*uint, error) {
	identifier := uint(len(app.contexts))
	return app.beginWithContext(identifier, name)
}

// Retrieve retrieves entry data from a context
func (app *application) Retrieve(identifier uint, retrieval delimiters.Delimiter) ([]byte, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		if pContext.pDataIndex == nil {
			str := fmt.Sprintf("the database for identifier (%d) contains no data", identifier)
			return nil, errors.New(str)
		}

		pointer, err := pContext.currentHeader.Fetch(retrieval)
		if err != nil {
			return nil, err
		}

		if pointer.IsDeleted() {
			return nil, errors.New("the requested retrieval has been deleted")
		}

		return app.readEntry(pContext.pFile, *pContext.pDataIndex, pointer.Delimiter())
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// RetrieveAll retrieves multiple entry data from context
func (app *application) RetrieveAll(identifier uint, retrievals delimiters.Delimiters) ([][]byte, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		return app.readEntries(pContext.pFile, *pContext.pDataIndex, retrievals)
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
			path:          pContext.path,
			name:          pContext.name,
			pDataIndex:    pContext.pDataIndex,
			insertions:    entries,
			currentHeader: pContext.currentHeader,
			deletions:     pContext.deletions,
			pFile:         pContext.pFile,
			pLock:         pContext.pLock,
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
			path:          pContext.path,
			name:          pContext.name,
			pDataIndex:    pContext.pDataIndex,
			insertions:    entries,
			currentHeader: pContext.currentHeader,
			deletions:     pContext.deletions,
			pFile:         pContext.pFile,
			pLock:         pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Delete deletes an entry from a context
func (app *application) Delete(identifier uint, delete delimiters.Delimiter) error {
	if pContext, ok := app.contexts[identifier]; ok {
		retDeletes, err := app.mergeDelete(pContext.deletions, []delimiters.Delimiter{
			delete,
		})

		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:          pContext.path,
			name:          pContext.name,
			pDataIndex:    pContext.pDataIndex,
			currentHeader: pContext.currentHeader,
			insertions:    pContext.insertions,
			deletions:     retDeletes,
			pFile:         pContext.pFile,
			pLock:         pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// DeleteAll deletes multiple entries from from a context
func (app *application) DeleteAll(identifier uint, deletes delimiters.Delimiters) error {
	if pContext, ok := app.contexts[identifier]; ok {
		retDeletes, err := app.mergeDelete(pContext.deletions, deletes.List())
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:          pContext.path,
			name:          pContext.name,
			pDataIndex:    pContext.pDataIndex,
			currentHeader: pContext.currentHeader,
			insertions:    pContext.insertions,
			deletions:     retDeletes,
			pFile:         pContext.pFile,
			pLock:         pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Commit commits a context
func (app *application) Commit(identifier uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		// create the temporary file name:
		value := strconv.Itoa(rand.Int())
		pHash, err := app.hashAdapter.FromBytes([]byte(value))
		if err != nil {
			return err
		}

		// create the destination path:
		destinationPath := filepath.Join(append(pContext.path[:len(pContext.path)-1], pHash.String())...)

		// create the temporary file:
		destinationFile, err := os.Create(destinationPath)
		if err != nil {
			return err
		}

		// close the file, then cleanup:
		defer destinationFile.Close()
		defer os.Remove(destinationPath)

		// update the header states:
		updatedStates, err := app.updateStates(pContext.currentHeader, pContext.insertions, pContext.deletions)
		if err != nil {
			return err
		}

		// update the header states on file:
		_, err = app.writeHeader(destinationFile, updatedStates)
		if err != nil {
			return err
		}

		// copy the existing data:
		dataIndex := uint64(0)
		if pContext.pDataIndex != nil {
			dataIndex = *pContext.pDataIndex
		}

		_, err = pContext.pFile.Seek(int64(dataIndex), io.SeekStart)
		if err != nil {
			return err
		}

		_, err = destinationFile.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}

		buffer := make([]byte, 1024)
		for {
			amountRead, err := pContext.pFile.Read(buffer)
			if err != nil && err != io.EOF {
				return err
			}

			if amountRead == 0 {
				break
			}

			amountWritten, err := destinationFile.Write(buffer[0:amountRead])
			if err != nil {
				return err
			}

			if amountRead != amountWritten {
				str := fmt.Sprintf("there was an error while copying data, amount bytes read: %d, amount bytes written: %d", amountRead, amountWritten)
				return errors.New(str)
			}
		}

		// write the insertions:
		if pContext.insertions != nil {
			err = app.writeInsertions(destinationFile, pContext.insertions)
			if err != nil {
				return err
			}
		}

		// replace the file:
		originPath := filepath.Join(pContext.path...)
		err = app.replaceFile(originPath, pContext.pFile, destinationFile)
		if err != nil {
			return err
		}

		// close the context and reopens it:
		err = app.Close(identifier)
		if err != nil {
			return err
		}

		_, err = app.beginWithContext(identifier, pContext.name)
		if err != nil {
			return err
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// DeleteState deletes a states from the context by state index
func (app *application) DeleteState(identifier uint, stateIndex uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		list := pContext.currentHeader.List()
		if len(list)-1 < int(stateIndex) {
			str := fmt.Sprintf("the header contains %d states, the requested state index (%d) does not exists", len(list), stateIndex)
			return errors.New(str)
		}

		currentState := list[stateIndex]
		if currentState.IsDeleted() {
			str := fmt.Sprintf("the state (index: %d) is already deleted", stateIndex)
			return errors.New(str)
		}

		stateBuilder := app.stateBuilder.Create().IsDeleted()
		if currentState.HasPointers() {
			pointers := currentState.Pointers()
			stateBuilder.WithPointers(pointers)
		}

		updatedState, err := stateBuilder.Now()
		if err != nil {
			return err
		}

		list[stateIndex] = updatedState
		updatedStates, err := app.statesBuilder.Create().WithList(list).Now()
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:          pContext.path,
			name:          pContext.name,
			pDataIndex:    pContext.pDataIndex,
			currentHeader: updatedStates,
			insertions:    pContext.insertions,
			deletions:     pContext.deletions,
			pFile:         pContext.pFile,
			pLock:         pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// RecoverState recovers a state in context by state index
func (app *application) RecoverState(identifier uint, stateIndex uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		list := pContext.currentHeader.List()
		if len(list)-1 < int(stateIndex) {
			str := fmt.Sprintf("the header contains %d states, the requested state index (%d) does not exists", len(list), stateIndex)
			return errors.New(str)
		}

		currentState := list[stateIndex]
		if !currentState.IsDeleted() {
			str := fmt.Sprintf("the state (index: %d) has not been deleted", stateIndex)
			return errors.New(str)
		}

		stateBuilder := app.stateBuilder.Create()
		if currentState.HasPointers() {
			pointers := currentState.Pointers()
			stateBuilder.WithPointers(pointers)
		}

		updatedState, err := stateBuilder.Now()
		if err != nil {
			return err
		}

		list[stateIndex] = updatedState
		updatedStates, err := app.statesBuilder.Create().WithList(list).Now()
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			path:          pContext.path,
			name:          pContext.name,
			pDataIndex:    pContext.pDataIndex,
			currentHeader: updatedStates,
			insertions:    pContext.insertions,
			deletions:     pContext.deletions,
			pFile:         pContext.pFile,
			pLock:         pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// DeletedStateIndexes returns the deleted state indexes
func (app *application) DeletedStateIndexes(identifier uint) ([]uint, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		indexes := []uint{}
		list := pContext.currentHeader.List()
		for idx, oneState := range list {
			if oneState.IsDeleted() {
				indexes = append(indexes, uint(idx))
			}
		}

		return indexes, nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// StatesAmount returns the amount of states
func (app *application) StatesAmount(identifier uint) (*uint, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		list := pContext.currentHeader.List()
		amount := uint(len(list))
		return &amount, nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// Purge purges the previous states and only keep the latest one.  It also deletes previously deleted entries
func (app *application) Purge(context uint) error {
	return nil
}

// Close closes a context
func (app *application) Close(identifier uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		err := pContext.pLock.Unlock()
		if err != nil {
			return err
		}

		err = pContext.pFile.Close()
		if err != nil {
			return err
		}

		delete(app.contexts, identifier)
		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

func (app *application) replaceFile(sourcePath string, pDestination *os.File, pSource *os.File) error {
	// Seek the destination file to the beginning:
	_, err := pSource.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = pDestination.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// copy the data:
	_, err = io.Copy(pDestination, pSource)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) writeHeader(file *os.File, header states.States) (*uint64, error) {
	bytes, err := app.statesAdapter.InstancesToBytes(header)
	if err != nil {
		return nil, err
	}

	length := len(bytes)
	lengthBytes := infra_bytes.Uint64ToBytes(uint64(length))
	output := append(lengthBytes, bytes...)

	// start at the beginning of the file:
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// write the header data:
	amountWritten, err := file.Write(output)
	if err != nil {
		return nil, err
	}

	if len(output) != amountWritten {
		str := fmt.Sprintf("expected to write %d length of data, %d actually written", len(output), amountWritten)
		return nil, errors.New(str)
	}

	outputLength := uint64(len(output))
	return &outputLength, nil
}

func (app *application) writeInsertions(file *os.File, insertions entries.Entries) error {
	list := insertions.List()

	// seek the end of the file:
	_, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	for _, oneInsertion := range list {
		data := oneInsertion.Bytes()
		amountWritten, err := file.Write(data)
		if err != nil {
			return err
		}

		if len(data) != amountWritten {
			str := fmt.Sprintf("expected to write %d length of data, %d actually written", len(data), amountWritten)
			return errors.New(str)
		}
	}

	return nil
}

func (app *application) beginWithContext(requestedContext uint, name string) (*uint, error) {
	fullPath := append(app.basepath, name)
	filePath := filepath.Join(fullPath...)

	var pFile *os.File
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		dir := filepath.Dir(filePath)
		err := os.MkdirAll(dir, os.ModePerm) // Create the directory path
		if err != nil {
			return nil, err
		}

		// Create the file
		pFile, err = os.Create(filePath)
		if err != nil {
			return nil, err
		}
	}

	// lock the origin file:
	originPath := filepath.Join(filePath)
	pLock := fslock.New(originPath)
	err := pLock.TryLock()
	if err != nil {
		str := fmt.Sprintf("failed to acquire lock: %s", err.Error())
		return nil, errors.New(str)
	}

	if pFile == nil {
		pOpenFile, err := os.OpenFile(filePath, os.O_RDWR, os.ModeAppend)
		if err != nil {
			str := fmt.Sprintf("failed to open file: %s", err.Error())
			return nil, errors.New(str)
		}

		pFile = pOpenFile
	}

	currentHeader, pDataIndex, err := app.readHeader(pFile)
	if err != nil {
		state, err := app.stateBuilder.Create().Now()
		if err != nil {
			return nil, err
		}

		states, err := app.statesBuilder.WithList([]states.State{
			state,
		}).Now()

		if err != nil {
			return nil, err
		}

		currentHeader = states
	}

	app.contexts[requestedContext] = &context{
		path:          fullPath,
		name:          name,
		pDataIndex:    pDataIndex,
		currentHeader: currentHeader,
		insertions:    nil,
		deletions:     nil,
		pFile:         pFile,
		pLock:         pLock,
	}

	return &requestedContext, nil
}

func (app *application) readEntries(file *os.File, dataIndex uint64, retrievals delimiters.Delimiters) ([][]byte, error) {
	output := [][]byte{}
	list := retrievals.List()
	for idx, oneRetrieval := range list {
		bytes, err := app.readEntry(file, dataIndex, oneRetrieval)
		if err != nil {
			str := fmt.Sprintf("could not read entry (pointer index: %d): %s", idx, err.Error())
			return nil, errors.New(str)
		}

		output = append(output, bytes)
	}

	return output, nil
}

func (app *application) readHeader(file *os.File) (states.States, *uint64, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, nil, errors.ErrUnsupported
	}

	// read the first int64 of the file:
	lengthBytes, err := app.readBytes(file, 0, amountOfBytesIntUint64)
	if err != nil {
		return nil, nil, err
	}

	// convert the bytes to the length:
	length := infra_bytes.BytesToUint64(lengthBytes)

	// read the data:
	headerBytes, err := app.readBytes(file, amountOfBytesIntUint64, int64(length))
	if err != nil {
		return nil, nil, err
	}

	retIns, _, err := app.statesAdapter.BytesToInstances(headerBytes)
	if err != nil {
		return nil, nil, err
	}

	offset := length + uint64(len(lengthBytes))
	return retIns, &offset, nil
}

func (app *application) readEntry(file *os.File, dataIndex uint64, retrieval delimiters.Delimiter) ([]byte, error) {
	index := dataIndex + retrieval.Index()
	length := retrieval.Length()
	return app.readBytes(file, int64(index), int64(length))
}

func (app *application) readBytes(file *os.File, index int64, length int64) ([]byte, error) {
	_, err := file.Seek(index, io.SeekStart)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, int(length))
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

func (app *application) mergeDelete(original delimiters.Delimiters, newEntries []delimiters.Delimiter) (delimiters.Delimiters, error) {
	list := []delimiters.Delimiter{}
	if original != nil {
		list = append(list, original.List()...)
	}

	list = append(list, newEntries...)
	return app.delimitersBuilder.Create().
		WithList(list).
		Now()
}

func (app *application) updateStates(statesIns states.States, insert entries.Entries, deletes delimiters.Delimiters) (states.States, error) {
	if insert == nil && deletes == nil {
		return statesIns, nil
	}

	updatedStates := statesIns.List()
	if deletes != nil {
		retStatesList, err := app.updateStatesWithDeletes(statesIns, deletes)
		if err != nil {
			return nil, err
		}

		updatedStates = retStatesList
	}

	if insert != nil {
		pointers, err := app.createPointers(insert)
		if err != nil {
			return nil, err
		}

		newState, err := app.stateBuilder.Create().WithPointers(pointers).Now()
		if err != nil {
			return nil, err
		}

		updatedStates = append(updatedStates, newState)
	}

	return app.statesBuilder.Create().
		WithList(updatedStates).
		Now()
}

func (app *application) updateStatesWithDeletes(statesIns states.States, deletes delimiters.Delimiters) ([]states.State, error) {
	updatedStates := []states.State{}
	list := statesIns.List()
	deletesList := deletes.List()
	for _, oneDelete := range deletesList {
		for _, oneState := range list {
			if oneState.IsDeleted() || !oneState.HasPointers() {
				updatedStates = append(updatedStates, oneState)
				continue
			}

			updatedPointersList := []pointers.Pointer{}
			pointersList := oneState.Pointers().List()
			for _, onePointer := range pointersList {
				if onePointer.IsDeleted() {
					updatedPointersList = append(updatedPointersList, onePointer)
					continue
				}

				delimiter := onePointer.Delimiter()
				if delimiter.Index() == oneDelete.Index() && delimiter.Length() == oneDelete.Length() {
					updatedPointer, err := app.pointerBuilder.Create().IsDeleted().WithDelimiter(delimiter).Now()
					if err != nil {
						return nil, err
					}

					updatedPointersList = append(updatedPointersList, updatedPointer)
					continue
				}

				updatedPointersList = append(updatedPointersList, onePointer)
			}

			var pointers pointers.Pointers
			if len(updatedPointersList) > 0 {
				updatedPointers, err := app.pointersBuilder.Create().WithList(updatedPointersList).Now()
				if err != nil {
					return nil, err
				}

				pointers = updatedPointers
			}

			stateBuilder := app.stateBuilder.Create()
			if pointers != nil {
				stateBuilder.WithPointers(pointers)
			}

			updatedState, err := stateBuilder.Now()
			if err != nil {
				return nil, err
			}

			updatedStates = append(updatedStates, updatedState)
		}
	}

	return updatedStates, nil
}

func (app *application) createPointers(entries entries.Entries) (pointers.Pointers, error) {
	pointersList := []pointers.Pointer{}
	list := entries.List()
	for _, oneEntry := range list {
		delimiter := oneEntry.Delimiter()
		pointer, err := app.pointerBuilder.Create().WithDelimiter(delimiter).Now()
		if err != nil {
			return nil, err
		}

		pointersList = append(pointersList, pointer)
	}

	return app.pointersBuilder.Create().
		WithList(pointersList).
		Now()
}
