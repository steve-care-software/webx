package tmp

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/engine/bytes/applications"
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
	"github.com/steve-care-software/webx/engine/bytes/domain/originals"
	infra_bytes "github.com/steve-care-software/webx/engine/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type application struct {
	namespaceAdapter  namespaces.Adapter
	namespacesBuilder namespaces.Builder
	namespaceBuilder  namespaces.NamespaceBuilder
	hashAdapter       hash.Adapter
	basepath          []string
	contexts          map[uint]*context
}

func createApplication(
	namespaceAdapter namespaces.Adapter,
	namespacesBuilder namespaces.Builder,
	namespaceBuilder namespaces.NamespaceBuilder,
	hashAdapter hash.Adapter,
	basepath []string,
) applications.Application {
	out := application{
		namespaceAdapter:  namespaceAdapter,
		namespacesBuilder: namespacesBuilder,
		namespaceBuilder:  namespaceBuilder,
		hashAdapter:       hashAdapter,
		basepath:          basepath,
		contexts:          map[uint]*context{},
	}

	return &out
}

// Begin begins a context on a database
func (app *application) Begin(name string) (*uint, error) {
	identifier := uint(len(app.contexts))
	return app.beginWithContext(identifier, name)
}

// Status returns the status, 0 = namespace
func (app *application) Status(context uint) (string, error) {
	if pContext, ok := app.contexts[context]; ok {
		return pContext.currentNamespace.Name(), nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, context)
	return "", errors.New(str)
}

// Namespace returns the namespace by name
func (app *application) Namespace(context uint, name string) (namespaces.Namespace, error) {
	if pContext, ok := app.contexts[context]; ok {
		return pContext.namespaces.Fetch(name)
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, context)
	return nil, errors.New(str)
}

// Namespaces returns the namespaces
func (app *application) Namespaces(context uint) (namespaces.Namespaces, error) {
	if pContext, ok := app.contexts[context]; ok {
		return pContext.namespaces, nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, context)
	return nil, errors.New(str)
}

// DeletedNamespaces returns the deleted namespaces
func (app *application) DeletedNamespaces(identifier uint) ([]string, error) {
	if pContext, ok := app.contexts[identifier]; ok {
		return pContext.namespaces.DeletedNames(), nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return nil, errors.New(str)
}

// SetNamespace sets the namespace
func (app *application) SetNamespace(identifier uint, name string) error {
	if pContext, ok := app.contexts[identifier]; ok {
		namespace, err := pContext.namespaces.Fetch(name)
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			dbPath:              pContext.dbPath,
			dbName:              pContext.dbName,
			namespaces:          pContext.namespaces,
			currentNamespace:    namespace,
			pNamespaceDataIndex: pContext.pNamespaceDataIndex,
			pFile:               pContext.pFile,
			pLock:               pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// InsertNamespace inserts a namespace
func (app *application) InsertNamespace(identifier uint, original originals.Original) error {
	if pContext, ok := app.contexts[identifier]; ok {
		name := original.Name()
		description := original.Description()
		newNamespace, err := app.createNamespace(
			name,
			description,
			false,
			nil,
		)

		if err != nil {
			return err
		}

		namespacesList := []namespaces.Namespace{}
		if pContext.namespaces != nil {
			namespacesList = pContext.namespaces.List()
		}

		namespacesList = append(namespacesList, newNamespace)
		namespaces, err := app.namespacesBuilder.Create().WithList(namespacesList).Now()
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			dbPath:              pContext.dbPath,
			dbName:              pContext.dbName,
			namespaces:          namespaces,
			currentNamespace:    pContext.currentNamespace,
			pNamespaceDataIndex: pContext.pNamespaceDataIndex,
			pFile:               pContext.pFile,
			pLock:               pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// OriginalNamespace updates a namespace
func (app *application) OriginalNamespace(identifier uint, original string, updated originals.Original) error {
	if pContext, ok := app.contexts[identifier]; ok {
		originalNamespace, err := pContext.namespaces.Fetch(original)
		if err != nil {
			return err
		}

		updatedNamespace, err := app.createNamespace(
			updated.Name(),
			updated.Description(),
			originalNamespace.IsDeleted(),
			originalNamespace.Iterations(),
		)

		if err != nil {
			return err
		}

		pIndex, err := pContext.namespaces.Index(original)
		if err != nil {
			return err
		}

		index := *pIndex
		namespacesList := pContext.namespaces.List()
		namespacesList = append(namespacesList[:index], namespacesList[index+1:]...) // remove the original namespace
		namespacesList = append(namespacesList, updatedNamespace)                    // add the new namespace
		namespaces, err := app.namespacesBuilder.Create().WithList(namespacesList).Now()
		if err != nil {
			return err
		}

		currentnamespace := pContext.currentNamespace
		if currentnamespace != nil {
			if currentnamespace.Name() == original {
				currentnamespace = updatedNamespace
			}
		}

		app.contexts[identifier] = &context{
			dbPath:              pContext.dbPath,
			dbName:              pContext.dbName,
			namespaces:          namespaces,
			currentNamespace:    currentnamespace,
			pNamespaceDataIndex: pContext.pNamespaceDataIndex,
			pFile:               pContext.pFile,
			pLock:               pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// DeleteNamespace deletes a namespace
func (app *application) DeleteNamespace(identifier uint, name string) error {
	if pContext, ok := app.contexts[identifier]; ok {
		if pContext.currentNamespace != nil {
			currentnamespace := pContext.currentNamespace
			if currentnamespace.Name() == name {
				str := fmt.Sprintf("the namespace (%s) cannot be deleted because it is the current namespace", name)
				return errors.New(str)
			}
		}

		retNamespace, err := pContext.namespaces.Fetch(name)
		if err != nil {
			return err
		}

		if retNamespace.IsDeleted() {
			str := fmt.Sprintf("the namespace (%s) has already been deleted", name)
			return errors.New(str)
		}

		deletedNamespace, err := app.createNamespace(
			retNamespace.Name(),
			retNamespace.Description(),
			true,
			retNamespace.Iterations(),
		)

		if err != nil {
			return err
		}

		pIndex, err := pContext.namespaces.Index(name)
		if err != nil {
			return err
		}

		index := *pIndex
		namespacesList := pContext.namespaces.List()
		namespacesList = append(namespacesList[:index], namespacesList[index+1:]...) // remove the old namespace
		namespacesList = append(namespacesList[:index], deletedNamespace)            // add the deleted namespace
		namespaces, err := app.namespacesBuilder.Create().WithList(namespacesList).Now()
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			dbPath:              pContext.dbPath,
			dbName:              pContext.dbName,
			namespaces:          namespaces,
			currentNamespace:    pContext.currentNamespace,
			pNamespaceDataIndex: pContext.pNamespaceDataIndex,
			pFile:               pContext.pFile,
			pLock:               pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// RecoverNamespace recovers namespaces
func (app *application) RecoverNamespace(identifier uint, name string) error {
	if pContext, ok := app.contexts[identifier]; ok {
		currentnamespace := pContext.currentNamespace
		currentName := currentnamespace.Name()
		if !currentnamespace.IsDeleted() {
			str := fmt.Sprintf("the namespace (%s) is already active", name)
			return errors.New(str)
		}

		recoveredNamespace, err := app.createNamespace(
			currentName,
			currentnamespace.Description(),
			true,
			currentnamespace.Iterations(),
		)

		if err != nil {
			return err
		}

		pIndex, err := pContext.namespaces.Index(name)
		if err != nil {
			return err
		}

		index := *pIndex
		namespacesList := pContext.namespaces.List()
		namespacesList = append(namespacesList[:index], namespacesList[index+1:]...) // remove the old namespace
		namespacesList = append(namespacesList[:index], recoveredNamespace)          // add the recovered namespace
		namespaces, err := app.namespacesBuilder.Create().WithList(namespacesList).Now()
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			dbPath:              pContext.dbPath,
			dbName:              pContext.dbName,
			namespaces:          namespaces,
			currentNamespace:    currentnamespace,
			pNamespaceDataIndex: pContext.pNamespaceDataIndex,
			pFile:               pContext.pFile,
			pLock:               pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// PurgeNamespace purges namespace by name
func (app *application) PurgeNamespace(identifier uint, name string) error {
	if pContext, ok := app.contexts[identifier]; ok {
		currentnamespace := pContext.currentNamespace
		if currentnamespace.Name() == name {
			str := fmt.Sprintf("the namespace (%s) cannot be purged because it is the current namespace", name)
			return errors.New(str)
		}

		pIndex, err := pContext.namespaces.Index(name)
		if err != nil {
			return err
		}

		index := *pIndex
		namespacesList := pContext.namespaces.List()
		namespacesList = append(namespacesList[:index], namespacesList[index+1:]...) // remove the namespace
		namespaces, err := app.namespacesBuilder.Create().WithList(namespacesList).Now()
		if err != nil {
			return err
		}

		app.contexts[identifier] = &context{
			dbPath:              pContext.dbPath,
			dbName:              pContext.dbName,
			namespaces:          namespaces,
			currentNamespace:    currentnamespace,
			pNamespaceDataIndex: pContext.pNamespaceDataIndex,
			pFile:               pContext.pFile,
			pLock:               pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// PurgeNamespaces purges all deleted namespaces
func (app *application) PurgeNamespaces(identifier uint) error {
	if pContext, ok := app.contexts[identifier]; ok {
		var namespaces namespaces.Namespaces
		activeList := pContext.namespaces.ActiveList()
		if len(activeList) > 0 {
			retNamespaces, err := app.namespacesBuilder.Create().WithList(activeList).Now()
			if err != nil {
				return err
			}

			namespaces = retNamespaces
		}

		app.contexts[identifier] = &context{
			dbPath:              pContext.dbPath,
			dbName:              pContext.dbName,
			namespaces:          namespaces,
			currentNamespace:    pContext.currentNamespace,
			pNamespaceDataIndex: pContext.pNamespaceDataIndex,
			pFile:               pContext.pFile,
			pLock:               pContext.pLock,
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

// Commit updates the database for the context
func (app *application) Commit(context uint) error {
	return app.commit(context, nil)
}

// Purge purges the database (deleted the deleted namespaes, states, branches and layers)
func (app *application) Purge(context uint) error {
	err := app.PurgeNamespaces(context)
	if err != nil {
		return err
	}

	return nil
}

// Cleanup cleanups the database (reove all unused data)
func (app *application) Cleanup(context uint) error {
	return nil
}

// Close closes the context
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

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
	return errors.New(str)
}

func (app *application) beginWithContext(requestedContext uint, dbName string) (*uint, error) {
	fullPath := append(app.basepath, dbName)
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

	currentNamespaces, pDataIndex, _ := app.readNamespaces(pFile)
	app.contexts[requestedContext] = &context{
		dbPath:              fullPath,
		dbName:              dbName,
		namespaces:          currentNamespaces,
		currentNamespace:    nil,
		pNamespaceDataIndex: pDataIndex,
		pFile:               pFile,
		pLock:               pLock,
	}

	return &requestedContext, nil
}

func (app *application) readNamespaces(pFile *os.File) (namespaces.Namespaces, *uint64, error) {
	_, err := pFile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, nil, err
	}

	// read the first int64 of the file:
	lengthBytes, err := app.readBytes(pFile, 0, infra_bytes.AmountOfBytesIntUint64)
	if err != nil {
		return nil, nil, err
	}

	// convert the bytes to the length:
	length := infra_bytes.BytesToUint64(lengthBytes)

	// read the data:
	namespaceBytes, err := app.readBytes(pFile, infra_bytes.AmountOfBytesIntUint64, int64(length))
	if err != nil {
		return nil, nil, err
	}

	retIns, err := app.namespaceAdapter.BytesToInstances(namespaceBytes)
	if err != nil {
		return nil, nil, err
	}

	offset := length + uint64(len(lengthBytes))
	return retIns, &offset, nil
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

func (app *application) writeNamespaces(pFile *os.File, namespaces namespaces.Namespaces) error {
	// seek the end of the file:
	_, err := pFile.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	data, err := app.namespaceAdapter.InstancesToBytes(namespaces)
	if err != nil {
		return err
	}

	dataSize := infra_bytes.Uint64ToBytes(uint64(len(data)))
	toWrite := append(dataSize, data...)
	amountWritten, err := pFile.Write(toWrite)
	if err != nil {
		return err
	}

	if len(toWrite) != amountWritten {
		str := fmt.Sprintf("expected to write %d length of data, %d actually written", len(toWrite), amountWritten)
		return errors.New(str)
	}

	return nil
}

func (app *application) commit(identifier uint, metaData delimiters.Delimiter) error {
	if pContext, ok := app.contexts[identifier]; ok {
		// create the temporary file name:
		value := strconv.Itoa(rand.Int())
		pHash, err := app.hashAdapter.FromBytes([]byte(value))
		if err != nil {
			return err
		}

		// create the destination path:
		destinationPath := filepath.Join(append(pContext.dbPath[:len(pContext.dbPath)-1], pHash.String())...)

		// create the temporary file:
		destinationFile, err := os.Create(destinationPath)
		if err != nil {
			return err
		}

		// close the file, then cleanup:
		defer destinationFile.Close()
		defer os.Remove(destinationPath)

		// update the namespaces on file:
		err = app.writeNamespaces(destinationFile, pContext.namespaces)
		if err != nil {
			return err
		}

		// copy the existing data:
		dataIndex := uint64(0)
		if pContext.pNamespaceDataIndex != nil {
			dataIndex = *pContext.pNamespaceDataIndex
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

		// replace the file:
		originPath := filepath.Join(pContext.dbPath...)
		err = app.replaceFile(originPath, pContext.pFile, destinationFile)
		if err != nil {
			return err
		}

		// close the context and reopens it:
		err = app.Close(identifier)
		if err != nil {
			return err
		}

		_, err = app.beginWithContext(identifier, pContext.dbName)
		if err != nil {
			return err
		}

		return nil
	}

	str := fmt.Sprintf(contextIdentifierUndefinedPattern, identifier)
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

func (app *application) createNamespace(
	name string,
	description string,
	isDeleted bool,
	iterations delimiters.Delimiter,
) (namespaces.Namespace, error) {
	builder := app.namespaceBuilder.Create().WithName(name).WithDescription(description)
	if isDeleted {
		builder.IsDeleted()
	}

	if iterations != nil {
		builder.WithIterations(iterations)
	}

	return builder.Now()
}
