package files

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/commits"
	"github.com/steve-care-software/webx/databases/domain/commits/histories"
	"github.com/steve-care-software/webx/databases/domain/configs"
	"github.com/steve-care-software/webx/databases/domain/connections"
	"github.com/steve-care-software/webx/databases/domain/connections/contents"
	commit_contents "github.com/steve-care-software/webx/databases/domain/contents/commits"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
	"github.com/steve-care-software/webx/databases/infrastructure/restapis/clients"
)

type application struct {
	clientApplicationBuilder    clients.Builder
	connectionsBuilder          connections.Builder
	connectionBuilder           connections.ConnectionBuilder
	contentsBuilder             contents.Builder
	contentBuilder              contents.ContentBuilder
	commitHistoriesAdapter      histories.Adapter
	commitHistoriesBuilder      histories.Builder
	commitBuilder               commits.Builder
	commitContentAdapter        commit_contents.Adapter
	commitContentBuilder        commit_contents.Builder
	referenceAdapter            references.Adapter
	referenceBuilder            references.Builder
	referenceContentKeysBuilder references.ContentKeysBuilder
	referenceContentKeyBuilder  references.ContentKeyBuilder
	referenceCommitsBuilder     references.CommitsBuilder
	referenceCommitBuilder      references.CommitBuilder
	referencePointerBuilder     references.PointerBuilder
	hashTreeBuilder             hashtrees.Builder
	dirPath                     string
	dstExtension                string
	bckExtension                string
	readChunkSize               uint
	contexts                    map[uint]*context
}

func createApplication(
	clientApplicationBuilder clients.Builder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	contentsBuilder contents.Builder,
	contentBuilder contents.ContentBuilder,
	commitHistoriesAdapter histories.Adapter,
	commitHistoriesBuilder histories.Builder,
	commitBuilder commits.Builder,
	commitContentAdapter commit_contents.Adapter,
	commitContentBuilder commit_contents.Builder,
	referenceAdapter references.Adapter,
	referenceBuilder references.Builder,
	referenceContentKeysBuilder references.ContentKeysBuilder,
	referenceContentKeyBuilder references.ContentKeyBuilder,
	referenceCommitsBuilder references.CommitsBuilder,
	referenceCommitBuilder references.CommitBuilder,
	referencePointerBuilder references.PointerBuilder,
	hashTreeBuilder hashtrees.Builder,
	dirPath string,
	dstExtension string,
	bckExtension string,
	readChunkSize uint,
) applications.Application {
	out := application{
		clientApplicationBuilder:    clientApplicationBuilder,
		connectionsBuilder:          connectionsBuilder,
		connectionBuilder:           connectionBuilder,
		contentsBuilder:             contentsBuilder,
		contentBuilder:              contentBuilder,
		commitHistoriesAdapter:      commitHistoriesAdapter,
		commitHistoriesBuilder:      commitHistoriesBuilder,
		commitBuilder:               commitBuilder,
		commitContentAdapter:        commitContentAdapter,
		commitContentBuilder:        commitContentBuilder,
		referenceAdapter:            referenceAdapter,
		referenceBuilder:            referenceBuilder,
		referenceContentKeysBuilder: referenceContentKeysBuilder,
		referenceContentKeyBuilder:  referenceContentKeyBuilder,
		referenceCommitsBuilder:     referenceCommitsBuilder,
		referenceCommitBuilder:      referenceCommitBuilder,
		referencePointerBuilder:     referencePointerBuilder,
		hashTreeBuilder:             hashTreeBuilder,
		dirPath:                     dirPath,
		dstExtension:                dstExtension,
		bckExtension:                bckExtension,
		readChunkSize:               readChunkSize,
		contexts:                    map[uint]*context{},
	}

	return &out
}

// Exists returns true if the database exists, false otherwise
func (app *application) Exists(name string) (bool, error) {
	path := filepath.Join(app.dirPath, name)
	fileInfo, err := os.Stat(path)
	if err == nil {
		return !fileInfo.IsDir(), nil
	}

	return false, nil
}

// New creates a new database
func (app *application) New(name string) error {
	if _, err := os.Stat(app.dirPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(app.dirPath, filePermission)
		if err != nil {
			return err
		}
	}

	path := filepath.Join(app.dirPath, name)
	_, err := os.Stat(path)
	if err == nil {
		str := fmt.Sprintf("the database (name: %s) already exists and therefore cannot be created again", name)
		return errors.New(str)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	return file.Close()
}

// Delete deletes an existing database
func (app *application) Delete(name string) error {
	path := filepath.Join(app.dirPath, name)
	pInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if pInfo.IsDir() {
		str := fmt.Sprintf("the name (%s) was expected to be a file, not a directory", name)
		return errors.New(str)
	}

	return os.Remove(path)
}

// Connections returns the active connections
func (app *application) Connections() (connections.Connections, error) {
	connectionsList := []connections.Connection{}
	for _, oneContext := range app.contexts {
		builder := app.connectionBuilder.Create().
			WithIdentifier(oneContext.identifier).
			WithName(oneContext.name)

		if len(oneContext.contentList) > 0 {
			contents, err := app.contentsBuilder.Create().
				WithList(oneContext.contentList).
				Now()

			if err != nil {
				return nil, err
			}

			builder.WithContents(contents)
		}

		if len(oneContext.peerList) > 0 {
			builder.WithPeers(oneContext.peerList)
		}

		connection, err := builder.Now()
		if err != nil {
			return nil, err
		}

		connectionsList = append(connectionsList, connection)
	}

	return app.connectionsBuilder.Create().
		WithList(connectionsList).
		Now()
}

// Open opens a context on a given database
func (app *application) Open(name string) (*uint, error) {
	reference, offset, err := app.retrieveReference(name)
	if err != nil {
		return nil, err
	}

	// open the connection:
	path := filepath.Join(app.dirPath, name)
	pConn, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// create a Lock instance on the path:
	pLock := fslock.New(path)

	// create the context:
	pContext := &context{
		identifier:  uint(len(app.contexts)),
		pConn:       pConn,
		pLock:       pLock,
		name:        name,
		reference:   reference,
		dataOffset:  offset,
		contentList: []contents.Content{},
		peerList:    []*url.URL{},
	}

	app.contexts[pContext.identifier] = pContext
	return &pContext.identifier, nil
}

func (app *application) retrieveReference(name string) (references.Reference, uint, error) {
	path := filepath.Join(app.dirPath, name)
	pConn, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}

	defer pConn.Close()

	// read the reference length in bytes:
	refLengthBytes := make([]byte, expectedReferenceBytesLength)
	refAmount, err := pConn.Read(refLengthBytes)
	if err != nil || refAmount > 0 {
		return nil, 0, nil
	}

	if refAmount != expectedReferenceBytesLength {
		str := fmt.Sprintf("%d bytes were expected to be read when reading the reference length bytes, %d actually read", expectedReferenceBytesLength, refAmount)
		return nil, 0, errors.New(str)
	}

	// convert the reference length to uint64:
	refLength := binary.LittleEndian.Uint64(refLengthBytes)

	// read the reference data:
	refAllBytes := []byte{}
	originalOffset := int64(refLength) + expectedReferenceBytesLength
	offset := originalOffset
	for {
		refContentBytes := make([]byte, app.readChunkSize)
		refContentAmount, err := pConn.ReadAt(refContentBytes, offset)
		if err != nil {
			return nil, 0, err
		}

		refAllBytes = append(refAllBytes, refContentBytes...)
		offset += int64(refContentAmount)
		if refContentAmount != int(refLength) {
			break
		}
	}

	// convert the content to a reference instance:
	ins, err := app.referenceAdapter.ToReference(refAllBytes)
	if err != nil {
		return nil, 0, err
	}

	return ins, uint(originalOffset), nil
}

// ContentKeysByKind returns the contentKeys by context and kind
func (app *application) ContentKeysByKind(context uint, kind uint) (references.ContentKeys, error) {
	contentKeys, err := app.contentKeys(context)
	if err != nil {
		return nil, err
	}

	list := contentKeys.ListByKind(kind)
	if len(list) <= 0 {
		str := fmt.Sprintf("there is no ContentKey instances for the given kind: %d", kind)
		return nil, errors.New(str)
	}

	return app.referenceContentKeysBuilder.Create().WithList(list).Now()
}

func (app *application) contentKeys(context uint) (references.ContentKeys, error) {
	if pContext, ok := app.contexts[context]; ok {
		if pContext.reference == nil {
			str := fmt.Sprintf("there is zero (0) ContentKey in the given context: %d", context)
			return nil, errors.New(str)
		}

		return pContext.reference.ContentKeys(), nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot return the Content instance", context)
	return nil, errors.New(str)
}

// CommitByHash returns the commit by hash
func (app *application) CommitByHash(context uint, hash hash.Hash) (commits.Commit, error) {
	commits, err := app.commits(context)
	if err != nil {
		return nil, err
	}

	refCommit, err := commits.Fetch(hash)
	if err != nil {
		return nil, err
	}

	return app.retrieveCommitByCommitReference(context, refCommit)
}

// Histories returns the commits histories on a context
func (app *application) Histories(context uint) (histories.Histories, error) {
	refCommits, err := app.commits(context)
	if err != nil {
		return nil, err
	}

	commitsList := []commits.Commit{}
	refCommitsList := refCommits.List()
	for _, oneRefCommit := range refCommitsList {
		commit, err := app.retrieveCommitByCommitReference(context, oneRefCommit)
		if err != nil {
			return nil, err
		}

		commitsList = append(commitsList, commit)
	}

	return app.commitHistoriesAdapter.FromCommitsToHistories(commitsList)
}

func (app *application) commits(context uint) (references.Commits, error) {
	if pContext, ok := app.contexts[context]; ok {
		if pContext.reference == nil {
			str := fmt.Sprintf("there is zero (0) Commit in the given context: %d", context)
			return nil, errors.New(str)
		}

		return pContext.reference.Commits(), nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot return the Commits instance", context)
	return nil, errors.New(str)
}

// Read reads a pointer on a context
func (app *application) Read(context uint, pointer references.Pointer) ([]byte, error) {
	if pContext, ok := app.contexts[context]; ok {
		offset := pContext.dataOffset + pointer.From()
		length := pointer.Length()
		contentBytes := make([]byte, length)
		refContentAmount, err := pContext.pConn.ReadAt(contentBytes, int64(offset))
		if err != nil {
			return nil, err
		}

		if refContentAmount != int(length) {
			str := fmt.Sprintf("the Read operation was expected to read %d bytes, %d returned", length, refContentAmount)
			return nil, errors.New(str)
		}

		return contentBytes, nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot Read using this context", context)
	return nil, errors.New(str)
}

// ReadByHash reads content by hash
func (app *application) ReadByHash(context uint, hash hash.Hash) ([]byte, error) {
	contentKey, err := app.retrieveActiveContentKeyByHash(context, hash)
	if err != nil {
		return nil, err
	}

	return app.Read(context, contentKey.Content())
}

func (app *application) retrieveActiveContentKeyByHash(context uint, hash hash.Hash) (references.ContentKey, error) {
	contentKeys, err := app.contentKeys(context)
	if err != nil {
		return nil, err
	}

	return contentKeys.Fetch(hash)
}

// ReadAll read pointers on a context
func (app *application) ReadAll(context uint, pointers []references.Pointer) ([][]byte, error) {
	output := [][]byte{}
	for _, onePointer := range pointers {
		content, err := app.Read(context, onePointer)
		if err != nil {
			return nil, err
		}

		output = append(output, content)
	}

	return output, nil
}

// ReadAllByHashes reads content by hashes
func (app *application) ReadAllByHashes(context uint, hashes []hash.Hash) ([][]byte, error) {
	output := [][]byte{}
	for _, oneHash := range hashes {
		content, err := app.ReadByHash(context, oneHash)
		if err != nil {
			return nil, err
		}

		output = append(output, content)
	}

	return output, nil
}

// Write writes data to a context
func (app *application) Write(context uint, hash hash.Hash, data []byte, kind uint) error {
	if pContext, ok := app.contexts[context]; ok {
		contentIns, err := app.contentBuilder.Create().WithHash(hash).WithData(data).WithKind(kind).Now()
		if err != nil {
			return err
		}

		pContext.contentList = append(pContext.contentList, contentIns)
		app.contexts[context] = pContext
		return nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be written to", context)
	return errors.New(str)
}

// Cancel cancels a context
func (app *application) Cancel(context uint) error {
	if pContext, ok := app.contexts[context]; ok {
		pContext.contentList = []contents.Content{}
		app.contexts[context] = pContext
		return nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be canceled", context)
	return errors.New(str)
}

// Commit commits a context
func (app *application) Commit(context uint) error {
	// update the reference:
	updatedReference, err := app.updateReference(context)
	if err != nil {
		return err
	}

	if pContext, ok := app.contexts[context]; ok {
		// update database on disk:
		pConn, pOffset, err := app.updateDatabaseOnDisk(pContext, updatedReference)
		if err != nil {
			return err
		}

		// update the file connection and reference:
		app.contexts[context].reference = updatedReference
		app.contexts[context].dataOffset = *pOffset
		app.contexts[context].pConn = pConn
		return nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be comitted", context)
	return errors.New(str)
}

func (app *application) updateReference(context uint) (references.Reference, error) {
	if pContext, ok := app.contexts[context]; ok {
		// find the latest commit:
		builder := app.commitBuilder.Create()
		if pContext.reference != nil {
			refCommit := pContext.reference.Commits().Latest()
			latestCommit, err := app.retrieveCommitByCommitReference(context, refCommit)
			if err != nil {
				return nil, err
			}

			builder.WithParent(latestCommit)
		}

		blocks := [][]byte{}
		for _, oneContent := range pContext.contentList {
			// add the hash in the blocks for the commit values:
			blocks = append(blocks, oneContent.Hash().Bytes())
		}

		values, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
		if err != nil {
			return nil, err
		}

		createdOn := time.Now().UTC()
		commit, err := builder.WithValues(values).CreatedOn(createdOn).Now()
		if err != nil {
			return nil, err
		}

		commitHash := commit.Hash()
		commitValues := commit.Values()
		commitContentBuilder := app.commitContentBuilder.Create().WithHash(commitHash).WithValues(commitValues)
		if commit.HasParent() {
			commitParent := commit.Parent().Hash()
			commitContentBuilder.WithParent(commitParent)
		}

		commitContent, err := commitContentBuilder.Now()
		if err != nil {
			return nil, err
		}

		commitBytes, err := app.commitContentAdapter.ToContent(commitContent)
		if err != nil {
			return nil, err
		}

		// build the pointer:
		commitFrom := int64(0)
		if pContext.reference != nil {
			commitFrom = pContext.reference.Next()
		}

		commitPointer, err := app.referencePointerBuilder.Create().From(uint(commitFrom)).WithLength(uint(len(commitBytes))).Now()
		if err != nil {
			return nil, err
		}

		// build the commit reference:
		refCommit, err := app.referenceCommitBuilder.Create().WithHash(commitHash).WithPointer(commitPointer).CreatedOn(createdOn).Now()
		if err != nil {
			return nil, err
		}

		// save the pointers in the commit references:
		commitsList := []references.Commit{}
		if pContext.reference != nil {
			commitsList = pContext.reference.Commits().List()
		}

		commitsList = append(commitsList, refCommit)
		commits, err := app.referenceCommitsBuilder.Create().WithList(commitsList).Now()
		if err != nil {
			return nil, err
		}

		// get the pending content list:
		contentKeysList := []references.ContentKey{}
		if pContext.reference != nil {
			contentKeysList = pContext.reference.ContentKeys().List()
		}

		// save all content:
		offset := commitFrom
		for _, oneContent := range pContext.contentList {
			// build the pointer:
			dataLength := int64(len(oneContent.Data()))
			contentKeyPointer, err := app.referencePointerBuilder.Create().From(uint(offset)).WithLength(uint(dataLength)).Now()
			if err != nil {
				return nil, err
			}

			// build the content key:
			contentKey, err := app.referenceContentKeyBuilder.Create().WithHash(oneContent.Hash()).WithKind(oneContent.Kind()).WithContent(contentKeyPointer).WithCommit(commitHash).Now()
			if err != nil {
				return nil, err
			}

			//save the content key to the list:
			contentKeysList = append(contentKeysList, contentKey)

			// update the offset:
			offset += dataLength
		}

		updatedContentKeys, err := app.referenceContentKeysBuilder.Create().WithList(contentKeysList).Now()
		if err != nil {
			return nil, err
		}

		currentList := []*url.URL{}
		if pContext.reference != nil {
			if pContext.reference.HasPeers() {
				currentList = pContext.reference.Peers()
			}
		}

		updatedPeers, err := app.mergePeers(currentList, pContext.peerList)
		if err != nil {
			return nil, err
		}

		refBuilder := app.referenceBuilder.Create().
			WithContentKeys(updatedContentKeys).
			WithCommits(commits)

		if len(updatedPeers) > 0 {
			refBuilder.WithPeers(updatedPeers)
		}

		return refBuilder.Now()
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be comitted", context)
	return nil, errors.New(str)
}

func (app *application) mergePeers(currentList []*url.URL, newList []*url.URL) ([]*url.URL, error) {
	if len(currentList) <= 0 {
		return newList, nil
	}

	if len(newList) <= 0 {
		return currentList, nil
	}

	peersMap := map[string]*url.URL{}
	for _, onePeer := range currentList {
		keyname := onePeer.String()
		peersMap[keyname] = onePeer
	}

	for _, onePeer := range newList {
		keyname := onePeer.String()
		peersMap[keyname] = onePeer
	}

	updatedList := []*url.URL{}
	for _, onePeer := range peersMap {
		updatedList = append(updatedList, onePeer)
	}

	return updatedList, nil
}

func (app *application) retrieveCommitByCommitReference(context uint, refCommit references.Commit) (commits.Commit, error) {
	pointer := refCommit.Pointer()
	contentBytes, err := app.Read(context, pointer)
	if err != nil {
		return nil, err
	}

	commitContent, err := app.commitContentAdapter.ToCommit(contentBytes)
	if err != nil {
		return nil, err
	}

	values := commitContent.Values()
	createdOn := refCommit.CreatedOn()
	builder := app.commitBuilder.Create().WithValues(values).CreatedOn(createdOn)
	if commitContent.HasParent() {
		pParentHash := commitContent.Parent()
		parent, err := app.CommitByHash(context, *pParentHash)
		if err != nil {
			return nil, err
		}

		builder.WithParent(parent)
	}

	return builder.Now()
}

func (app *application) updateDatabaseOnDisk(context *context, updatedReference references.Reference) (*os.File, *uint, error) {
	// create a lock on the file:
	err := context.pLock.TryLock()
	if err != nil {
		return nil, nil, err
	}

	// release the lock on closing the method:
	defer context.pLock.Unlock()

	// write destination file:
	pOffset, err := app.writeDestinationOnDisk(context, updatedReference)
	if err != nil {
		return nil, nil, err
	}

	// create the source path:
	sourcePath := filepath.Join(app.dirPath, context.name)

	// create the backup path:
	backupFile := fmt.Sprintf("%s%s%s", context.name, fileNameExtensionDelimiter, app.bckExtension)
	backupPath := filepath.Join(app.dirPath, backupFile)

	// rename the source database to a backup file:
	backupPtr, err := os.Create(backupPath)
	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(backupPtr, context.pConn)
	if err != nil {
		return nil, nil, err
	}

	// close the backup file:
	err = backupPtr.Close()
	if err != nil {
		return nil, nil, err
	}

	// close the source connection:
	err = context.pConn.Close()
	if err != nil {
		return nil, nil, err
	}

	// delete the source database:
	err = os.Remove(sourcePath)
	if err != nil {
		return nil, nil, err
	}

	// rename the destination database to source:
	destinationFile := fmt.Sprintf("%s%s%s", context.name, fileNameExtensionDelimiter, app.dstExtension)
	destinationPath := filepath.Join(app.dirPath, destinationFile)
	err = os.Rename(destinationPath, sourcePath)
	if err != nil {
		return nil, nil, err
	}

	// delete the backup file:
	err = os.Remove(backupPath)
	if err != nil {
		return nil, nil, err
	}

	// re-open the source connection:
	pNewConn, err := os.Open(sourcePath)
	if err != nil {
		return nil, nil, err
	}

	return pNewConn, pOffset, nil
}

func (app *application) writeDestinationOnDisk(context *context, updatedReference references.Reference) (*uint, error) {
	// destination path:
	destinationFile := fmt.Sprintf("%s%s%s", context.name, fileNameExtensionDelimiter, app.dstExtension)
	destinationPath := filepath.Join(app.dirPath, destinationFile)

	// create the destination file:
	destination, err := os.Create(destinationPath)
	if err != nil {
		return nil, err
	}

	// close the destination:
	defer destination.Close()

	// convert the updated reference to data:
	refData, err := app.referenceToContent(updatedReference)
	if err != nil {
		return nil, err
	}

	// write the reference data on disk:
	writtenAmount, err := destination.Write(refData)
	if err != nil {
		return nil, err
	}

	if writtenAmount != len(refData) {
		str := fmt.Sprintf("%d bytes were expected to be writte while writing the updated reference bytes, %d actually written", len(refData), writtenAmount)
		return nil, errors.New(str)
	}

	offset := int64(context.dataOffset) + int64(len(refData)) + expectedReferenceBytesLength
	dataOffset := offset
	for {
		// read the file at offset:
		contentBytes := make([]byte, app.readChunkSize)
		amountRead, err := context.pConn.ReadAt(contentBytes, int64(dataOffset))
		if err != nil {
			break
		}

		if app.readChunkSize != uint(amountRead) {
			str := fmt.Sprintf("%d bytes were expected to be read from source database, %d actually read", app.readChunkSize, amountRead)
			return nil, errors.New(str)
		}

		// write content on destination:
		err = app.saveDataOnDisk(offset, contentBytes, destination)
		if err != nil {
			break
		}

		//update the offset:
		dataOffset += int64(amountRead)
	}

	// write the data on disk:
	for _, oneContent := range context.contentList {
		contentBytes := oneContent.Data()
		err = app.saveDataOnDisk(dataOffset, contentBytes, destination)
		if err != nil {
			break
		}

		// update the offset:
		dataOffset += int64(len(contentBytes))
	}

	retOffset := uint(offset)
	return &retOffset, nil
}

func (app *application) referenceToContent(reference references.Reference) ([]byte, error) {
	contentBytes, err := app.referenceAdapter.ToContent(reference)
	if err != nil {
		return nil, err
	}

	bytesLength := make([]byte, expectedReferenceBytesLength)
	binary.LittleEndian.PutUint64(bytesLength, uint64(len(contentBytes)))

	data := []byte{}
	data = append(data, bytesLength...)
	return append(data, contentBytes...), nil
}

func (app *application) saveDataOnDisk(offset int64, data []byte, pConn *os.File) error {
	// seek the file at the from byte:
	seekOffset, err := pConn.Seek(offset, 0)
	if err != nil {
		return err
	}

	if seekOffset != offset {
		str := fmt.Sprintf("the offset was expected to be %d, %d returned after file seek", offset, seekOffset)
		return errors.New(str)
	}

	// write the data on disk:
	amountWritten, err := pConn.Write(data)
	if err != nil {
		return err
	}

	amountExpected := len(data)
	if amountExpected != amountWritten {
		str := fmt.Sprintf("%d bytes were expected to be written, %d actually written", amountExpected, amountWritten)
		return errors.New(str)
	}

	return nil
}

// Share shares the database interactions with a new peer, using the given context
func (app *application) Share(context uint, peer *url.URL) error {
	if pContext, ok := app.contexts[context]; ok {
		pContext.peerList = append(pContext.peerList, peer)
		return nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be used to share to a peer", context)
	return errors.New(str)
}

// Push retrieves the commits from peers, then chain our commits to them using the given configuration
func (app *application) Push(name string, config configs.Config) error {
	// retrieve the reference from our database:
	reference, _, err := app.retrieveReference(name)
	if err != nil {
		return err
	}

	// open a context:
	pCurrentContext, err := app.Open(name)
	if err != nil {
		return err
	}

	if currentContext, ok := app.contexts[*pCurrentContext]; ok {
		// fetch the current history:
		currentHistories, err := app.Histories(currentContext.identifier)
		if err != nil {
			return err
		}

		// loop in the peers, if any:
		if reference.HasPeers() {
			betterCommits := map[string]commits.Commit{}
			peersList := reference.Peers()
			for _, onePeer := range peersList {
				// build the client application:
				clientApp, err := app.clientApplicationBuilder.Create().WithServer(onePeer).Now()
				if err != nil {
					return err
				}

				// open a context:
				pContext, err := clientApp.Open(name)
				if err != nil {
					return err
				}

				// download the histories of our peer:
				retHistories, err := clientApp.Histories(*pContext)
				if err != nil {
					return err
				}

				// compare with our current history:
				toDownloadHistoryList, err := currentHistories.Compare(retHistories)
				if err != nil {
					return err
				}

				if len(toDownloadHistoryList) < 0 {
					continue
				}

				// download the better commits, if needed:
				for _, oneHistory := range toDownloadHistoryList {
					commitHash := oneHistory.Commit()
					keyname := commitHash.String()
					if _, ok := betterCommits[keyname]; ok {
						continue
					}

					commit, err := clientApp.CommitByHash(*pContext, commitHash)
					if err != nil {
						return err
					}

					betterCommits[keyname] = commit
				}

			}
		}
	}

	str := fmt.Sprintf("the context (%d) was expected to be opened while executing the Push method", *pCurrentContext)
	return errors.New(str)
}

// Close closes a context
func (app *application) Close(context uint) error {
	if pContext, ok := app.contexts[context]; ok {
		err := pContext.pConn.Close()
		if err != nil {
			return err
		}

		delete(app.contexts, context)
		return nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be closed", context)
	return errors.New(str)
}
