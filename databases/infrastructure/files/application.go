package files

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/juju/fslock"
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/commits"
	"github.com/steve-care-software/webx/databases/domain/configs"
	commit_contents "github.com/steve-care-software/webx/databases/domain/contents/commits"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

type application struct {
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

// New creates a new database
func (app *application) New(name string) error {
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

func (app *application) saveReferenceOnDisk(name string, reference references.Reference) error {
	path := filepath.Join(app.dirPath, name)
	data, err := app.referenceToContent(reference)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, filePermission)
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

// Delete deletes an existing database
func (app *application) Delete(name string) error {
	path := filepath.Join(app.dirPath, name)
	pInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !pInfo.IsDir() {
		str := fmt.Sprintf("the name (%s) was expected to be a file, not a directory", name)
		return errors.New(str)
	}

	return os.Remove(path)
}

// Open opens a context at height, height is -1 if the head is requested
func (app *application) Open(name string, height int) (*uint, error) {
	path := filepath.Join(app.dirPath, name)
	pConn, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// read the reference length in bytes:
	refLengthBytes := make([]byte, expectedReferenceBytesLength)
	refAmount, err := pConn.Read(refLengthBytes)
	if err != nil {
		return nil, err
	}

	refLength := 0
	var reference references.Reference
	if refAmount > 0 {
		if refAmount != expectedReferenceBytesLength {
			str := fmt.Sprintf("%d bytes were expected to be read when reading the reference length bytes, %d actually read", expectedReferenceBytesLength, refAmount)
			return nil, errors.New(str)
		}

		// convert the reference length to uint64:
		refLength := binary.LittleEndian.Uint64(refLengthBytes)

		// read the reference data:
		refAllBytes := []byte{}
		offset := int64(refLength)
		for {
			refContentBytes := make([]byte, app.readChunkSize)
			refContentAmount, err := pConn.ReadAt(refContentBytes, offset)
			if err != nil {
				return nil, err
			}

			refAllBytes = append(refAllBytes, refContentBytes...)
			offset += int64(refContentAmount)
			if refContentAmount != int(refLength) {
				break
			}
		}

		// convert the content to a reference instance:
		refIns, err := app.referenceAdapter.ToReference(refAllBytes)
		if err != nil {
			return nil, err
		}

		reference = refIns
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
		dataOffset:  uint(refLength),
		contentList: []*content{},
	}

	return &pContext.identifier, nil
}

// ContentKeys returns the contentKeys by context
func (app *application) ContentKeys(context uint) (references.ContentKeys, error) {
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

// Commits returns the commits on a context
func (app *application) Commits(context uint) (references.Commits, error) {
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
	contentKeys, err := app.ContentKeys(context)
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
		pContext.contentList = append(pContext.contentList, &content{
			hash: hash,
			data: data,
			kind: kind,
		})

		app.contexts[context] = pContext
		return nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be written to", context)
	return errors.New(str)
}

// Cancel cancels a context
func (app *application) Cancel(context uint) error {
	if pContext, ok := app.contexts[context]; ok {
		pContext.contentList = []*content{}
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
		err := app.updateDatabaseOnDisk(pContext, updatedReference)
		if err != nil {
			return err
		}

		// delete the context:
		delete(app.contexts, context)
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
			blocks = append(blocks, oneContent.hash.Bytes())
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
		commitFrom := pContext.reference.Next()
		commitPointer, err := app.referencePointerBuilder.Create().From(uint(commitFrom)).WithLength(uint(len(commitBytes))).Now()
		if err != nil {
			return nil, err
		}

		// build the commit reference:
		refCommit, err := app.referenceCommitBuilder.Create().WithHash(commitHash).WithPointer(commitPointer).Now()
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
			dataLength := int64(len(oneContent.data))
			contentKeyPointer, err := app.referencePointerBuilder.Create().From(uint(offset)).WithLength(uint(dataLength)).Now()
			if err != nil {
				return nil, err
			}

			// build the content key:
			contentKey, err := app.referenceContentKeyBuilder.Create().WithHash(oneContent.hash).WithKind(oneContent.kind).WithContent(contentKeyPointer).WithCommit(commitHash).Now()
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

		return app.referenceBuilder.Create().
			WithContentKeys(updatedContentKeys).
			WithCommits(commits).
			Now()
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be comitted", context)
	return nil, errors.New(str)
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
		parentCommit, err := app.retrieveCommitByHash(context, *pParentHash)
		if err != nil {
			return nil, err
		}

		parent, err := app.retrieveCommitByCommitReference(context, parentCommit)
		if err != nil {
			return nil, err
		}

		builder.WithParent(parent)
	}

	return builder.Now()
}

func (app *application) retrieveCommitByHash(context uint, hash hash.Hash) (references.Commit, error) {
	commit, err := app.Commits(context)
	if err != nil {
		return nil, err
	}

	return commit.Fetch(hash)
}

func (app *application) updateDatabaseOnDisk(context *context, updatedReference references.Reference) error {
	// create a lock on the file:
	err := context.pLock.TryLock()
	if err != nil {
		return err
	}

	// release the lock on closing the method:
	defer context.pLock.Unlock()

	// write destination file:
	err = app.writeDestinationOnDisk(context, updatedReference)
	if err != nil {
		return err
	}

	// rename the source database to a backup file:
	sourcePath := filepath.Join(app.dirPath, context.name)
	backupPath := filepath.Join(sourcePath, app.bckExtension)
	backupPtr, err := os.Create(backupPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(backupPtr, context.pConn)
	if err != nil {
		return err
	}

	// close the backup file:
	err = backupPtr.Close()
	if err != nil {
		return err
	}

	// close the source connection:
	err = context.pConn.Close()
	if err != nil {
		return err
	}

	// delete the source database:
	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	// rename the destination database to source:
	destinationPath := filepath.Join(sourcePath, app.dstExtension)
	err = os.Rename(destinationPath, sourcePath)
	if err != nil {
		return err
	}

	// delete the backup file:
	return os.Remove(backupPath)
}

func (app *application) writeDestinationOnDisk(context *context, updatedReference references.Reference) error {
	// build the path:
	sourcePath := filepath.Join(app.dirPath, context.name)
	destinationPath := filepath.Join(sourcePath, app.dstExtension)

	// create the destination file:
	destination, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	// close the destination:
	defer destination.Close()

	// convert the updated reference to data:
	refData, err := app.referenceToContent(updatedReference)
	if err != nil {
		return err
	}

	// write the reference data on disk:
	writtenAmount, err := destination.Write(refData)
	if err != nil {
		return err
	}

	if writtenAmount != len(refData) {
		str := fmt.Sprintf("%d bytes were expected to be writte while writing the updated reference bytes, %d actually written", len(refData), writtenAmount)
		return errors.New(str)
	}

	offset := int64(context.dataOffset)
	for {
		// read the file at offset:
		contentBytes := make([]byte, app.readChunkSize)
		amountRead, err := context.pConn.ReadAt(contentBytes, int64(offset))
		if err != nil {
			return err
		}

		if app.readChunkSize != uint(amountRead) {
			str := fmt.Sprintf("%d bytes were expected to be read from source database, %d actually read", app.readChunkSize, amountRead)
			return errors.New(str)
		}

		// write content on destination:
		err = app.saveDataOnDisk(offset, contentBytes, destination)
		if err != nil {
			break
		}

		//update the offset:
		offset += int64(amountRead)
	}

	return nil
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

// Push retrieves the commits from peers, then chain our commits to them using the given configuration
func (app *application) Push(config configs.Config) error {
	return nil
}

// Close closes a context
func (app *application) Close(context uint) error {
	if pContext, ok := app.contexts[context]; ok {
		err := pContext.pConn.Close()
		if err != nil {
			return err
		}

		delete(app.contexts, context)
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be closed", context)
	return errors.New(str)
}
