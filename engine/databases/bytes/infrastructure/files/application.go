package files

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/steve-care-software/webx/engine/databases/bytes/applications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/listers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/modifications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
	infra_bytes "github.com/steve-care-software/webx/engine/databases/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type application struct {
	hashAdapter         hash.Adapter
	statesAdapter       states.Adapter
	statesBuilder       states.Builder
	stateBuilder        states.StateBuilder
	containersBuilder   containers.Builder
	containerBuilder    containers.ContainerBuilder
	pointersBuilder     pointers.Builder
	pointerBuilder      pointers.PointerBuilder
	modificationBuilder modifications.Builder
	entriesBuilder      entries.Builder
	deletesBuilder      deletes.Builder
	retrievalsBuilder   retrievals.Builder
	basepath            []string
	contexts            map[uint]*context
}

func createApplication(
	hashAdapter hash.Adapter,
	statesAdapter states.Adapter,
	statesBuilder states.Builder,
	stateBuilder states.StateBuilder,
	containersBuilder containers.Builder,
	containerBuilder containers.ContainerBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	modificationBuilder modifications.Builder,
	entriesBuilder entries.Builder,
	deletesBuilder deletes.Builder,
	retrievalsBuilder retrievals.Builder,
	basepath []string,
) applications.Application {
	out := application{
		hashAdapter:         hashAdapter,
		statesAdapter:       statesAdapter,
		statesBuilder:       statesBuilder,
		stateBuilder:        stateBuilder,
		containersBuilder:   containersBuilder,
		containerBuilder:    containerBuilder,
		pointersBuilder:     pointersBuilder,
		pointerBuilder:      pointerBuilder,
		modificationBuilder: modificationBuilder,
		entriesBuilder:      entriesBuilder,
		deletesBuilder:      deletesBuilder,
		retrievalsBuilder:   retrievalsBuilder,
		basepath:            basepath,
		contexts:            map[uint]*context{},
	}

	return &out
}

// Begin begins a context
func (app *application) Begin(name string) (*uint, error) {
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

	if pFile == nil {
		pOpenFile, err := os.Open(filePath)
		if err != nil {
			str := fmt.Sprintf("failed to open file: %s", err.Error())
			return nil, errors.New(str)
		}

		pFile = pOpenFile
	}

	currentHeader, err := app.readHeader(pFile)
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

	identifier := uint(len(app.contexts))
	app.contexts[identifier] = &context{
		path:          fullPath,
		currentHeader: currentHeader,
		insertions:    nil,
		deletions:     nil,
		pFile:         pFile,
	}

	return &identifier, nil
}

// List returns the list of pointers
func (app *application) List(identifier uint, lister listers.Lister) (retrievals.Retrievals, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		keyname := lister.Keyname()
		retrieval := lister.Retrieval()
		index := retrieval.Index()
		length := retrieval.Length()
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
			path:          pContext.path,
			insertions:    entries,
			currentHeader: pContext.currentHeader,
			deletions:     pContext.deletions,
			pFile:         pContext.pFile,
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
			insertions:    entries,
			currentHeader: pContext.currentHeader,
			deletions:     pContext.deletions,
			pFile:         pContext.pFile,
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
			path:          pContext.path,
			currentHeader: pContext.currentHeader,
			insertions:    pContext.insertions,
			deletions:     retDeletes,
			pFile:         pContext.pFile,
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
			path:          pContext.path,
			currentHeader: pContext.currentHeader,
			insertions:    pContext.insertions,
			deletions:     retDeletes,
			pFile:         pContext.pFile,
		}

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Commit commits a context
func (app *application) Commit(identifier uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		// lock the origin file:
		/*originPath := filepath.Join(pContext.path...)
		pLock := fslock.New(originPath)
		err := pLock.TryLock()
		if err != nil {
			str := fmt.Sprintf("failed to acquire lock: %s", err.Error())
			return errors.New(str)
		}

		defer pLock.Lock()*/
		defer pContext.pFile.Close()

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

		// close the file, then clearnup
		defer destinationFile.Close()
		//defer os.Remove(destinationPath)

		// copy the database to a temp database file:
		/*_, err = io.Copy(destinationFile, pContext.pFile)
		if err != nil {
			return err
		}*/

		// update the header states:
		updatedStates, err := app.insertInStates(pContext.currentHeader, pContext.insertions)
		if err != nil {
			return err
		}

		// update the header states on file:
		err = app.writeHeader(destinationFile, updatedStates)
		if err != nil {
			return err
		}

		// write the insertions:
		/*if pContext.insertions != nil {
			err = app.writeInsertions(destinationFile, pContext.insertions)
			if err != nil {
				return err
			}
		}*/

		// write the deletions:
		/*if pContext.deletions != nil {
			err = app.writeDeletions(destinationFile, pContext.deletions)
			if err != nil {
				return err
			}
		}*/

		// write the destination file back to the original file:
		/*_, err = io.Copy(pContext.pFile, destinationFile)
		if err != nil {
			return err
		}*/

		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
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

		delete(app.contexts, identifier)
		return nil
	}

	str := fmt.Sprintf(contentIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

func (app *application) writeHeader(file *os.File, header states.States) error {
	bytes, err := app.statesAdapter.InstancesToBytes(header)
	if err != nil {
		return err
	}

	length := len(bytes)
	lengthBytes := infra_bytes.Uint64ToBytes(uint64(length))
	output := append(lengthBytes, bytes...)

	// start at the beginning of the file:
	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.Write(output)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) writeInsertions(file *os.File, insertions entries.Entries) error {
	return nil
}

func (app *application) writeDeletions(file *os.File, deletions deletes.Deletes) error {
	return nil
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

func (app *application) readHeader(file *os.File) (states.States, error) {
	// read the first int64 of the file:
	lengthBytes, err := app.readBytes(file, 0, amountOfBytesIntUint64)
	if err != nil {
		return nil, err
	}

	// convert the bytes to the length:
	length := int64(infra_bytes.BytesToUint64(lengthBytes))

	// read the data:
	headerBytes, err := app.readBytes(file, amountOfBytesIntUint64, length)
	if err != nil {
		return nil, err
	}

	retIns, _, err := app.statesAdapter.BytesToInstances(headerBytes)
	if err != nil {
		return nil, err
	}

	return retIns, nil
}

func (app *application) readEntry(file *os.File, retrieval retrievals.Retrieval) ([]byte, error) {
	index := retrieval.Index()
	length := retrieval.Length()
	return app.readBytes(file, int64(index), int64(length))
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

func (app *application) insertInStates(statesIns states.States, entries entries.Entries) (states.States, error) {
	output := []states.State{}
	list := statesIns.List()
	length := len(list)
	for i := 0; i < length; i++ {
		index := length - 1 - i
		currentState := list[index]
		if currentState.IsDeleted() {
			output = append(output, currentState)
			continue
		}

		containers := currentState.Containers()
		newContainersList, err := app.createContainersForEntities(containers, entries)
		if err != nil {
			return nil, err
		}

		if !currentState.HasContainers() {
			currentContainers, err := app.containersBuilder.Create().
				WithList(newContainersList).
				Now()

			if err != nil {
				return nil, err
			}

			containers = currentContainers
		}

		updatedAfterInsertContainers, err := app.insertInContainers(containers, entries)
		if err != nil {
			return nil, err
		}

		updatedState, err := app.stateBuilder.Create().WithContainers(updatedAfterInsertContainers).Now()
		if err != nil {
			return nil, err
		}

		output = append(output, updatedState)
	}

	return app.statesBuilder.Create().
		WithList(output).
		Now()
}

func (app *application) createContainersForEntities(containersIns containers.Containers, entries entries.Entries) ([]containers.Container, error) {
	mp := map[string][]pointers.Pointer{}
	list := entries.List()
	for _, oneEntry := range list {
		keyname := oneEntry.Keyname()
		if containersIns != nil {
			_, err := containersIns.Fetch(keyname)
			if err == nil {
				continue
			}
		}

		delimiter := oneEntry.Delimiter()
		pointer, err := app.pointerBuilder.Create().WithDelimiter(delimiter).Now()
		if err != nil {
			return nil, err
		}

		if _, ok := mp[keyname]; !ok {
			mp[keyname] = []pointers.Pointer{}
		}

		mp[keyname] = append(mp[keyname], pointer)
	}

	output := []containers.Container{}
	for keyname, pointersList := range mp {
		pointers, err := app.pointersBuilder.Create().WithList(pointersList).Now()
		if err != nil {
			return nil, err
		}

		container, err := app.containerBuilder.Create().WithKeyname(keyname).WithPointers(pointers).Now()
		if err != nil {
			return nil, err
		}

		output = append(output, container)
	}

	return output, nil
}

func (app *application) insertInContainers(containersIns containers.Containers, entries entries.Entries) (containers.Containers, error) {
	updated := []containers.Container{}
	list := containersIns.List()
	for _, oneContainer := range list {
		updatedContainer, err := app.insertInContainer(oneContainer, entries)
		if err != nil {
			return nil, err
		}

		updated = append(updated, updatedContainer)
	}

	return app.containersBuilder.Create().
		WithList(updated).
		Now()
}

func (app *application) insertInContainer(container containers.Container, entries entries.Entries) (containers.Container, error) {
	pointersList := []pointers.Pointer{}
	keyname := container.Keyname()
	list := entries.List()
	for _, oneEntry := range list {
		if oneEntry.Keyname() != keyname {
			continue
		}

		delimiter := oneEntry.Delimiter()
		pointer, err := app.pointerBuilder.Create().WithDelimiter(delimiter).Now()
		if err != nil {
			return nil, err
		}

		pointersList = append(pointersList, pointer)
	}

	currentPointers := container.Pointers().List()
	updatedPointers, err := app.pointersBuilder.Create().
		WithList(append(currentPointers, pointersList...)).
		Now()

	if err != nil {
		return nil, err
	}

	return app.containerBuilder.Create().
		WithKeyname(keyname).
		WithPointers(updatedPointers).
		Now()
}
