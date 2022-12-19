package files

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

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
	referenceContentFactory     references.ContentFactory
	referenceContentBuilder     references.ContentBuilder
	referenceContentKeysBuilder references.ContentKeysBuilder
	referenceContentKeyBuilder  references.ContentKeyBuilder
	referenceCommitsBuilder     references.CommitsBuilder
	referenceCommitBuilder      references.CommitBuilder
	referencePointerBuilder     references.PointerBuilder
	hashTreeBuilder             hashtrees.Builder
	dirPath                     string
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
	referenceContentFactory references.ContentFactory,
	referenceContentBuilder references.ContentBuilder,
	referenceContentKeysBuilder references.ContentKeysBuilder,
	referenceContentKeyBuilder references.ContentKeyBuilder,
	referenceCommitsBuilder references.CommitsBuilder,
	referenceCommitBuilder references.CommitBuilder,
	referencePointerBuilder references.PointerBuilder,
	hashTreeBuilder hashtrees.Builder,
	dirPath string,
	bckExtension string,
	readChunkSize uint,
) applications.Application {
	out := application{
		commitBuilder:               commitBuilder,
		commitContentAdapter:        commitContentAdapter,
		commitContentBuilder:        commitContentBuilder,
		referenceAdapter:            referenceAdapter,
		referenceBuilder:            referenceBuilder,
		referenceContentFactory:     referenceContentFactory,
		referenceContentBuilder:     referenceContentBuilder,
		referenceContentKeysBuilder: referenceContentKeysBuilder,
		referenceContentKeyBuilder:  referenceContentKeyBuilder,
		referenceCommitsBuilder:     referenceCommitsBuilder,
		referenceCommitBuilder:      referenceCommitBuilder,
		referencePointerBuilder:     referencePointerBuilder,
		hashTreeBuilder:             hashTreeBuilder,
		dirPath:                     dirPath,
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

	content, err := app.referenceContentFactory.Create()
	if err != nil {
		return err
	}

	reference, err := app.referenceBuilder.Create().WithContent(content).Now()
	if err != nil {
		return err
	}

	return app.saveReferenceOnDisk(name, reference)
}

func (app *application) saveReferenceOnDisk(name string, reference references.Reference) error {
	path := filepath.Join(app.dirPath, name)
	contentBytes, err := app.referenceAdapter.ToContent(reference)
	if err != nil {
		return err
	}

	bytesLength := make([]byte, expectedReferenceBytesLength)
	binary.LittleEndian.PutUint64(bytesLength, uint64(len(contentBytes)))

	data := []byte{}
	data = append(data, bytesLength...)
	data = append(data, contentBytes...)
	return ioutil.WriteFile(path, data, filePermission)
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
	reference, err := app.referenceAdapter.ToReference(refAllBytes)
	if err != nil {
		return nil, err
	}

	pContext := &context{
		identifier:  uint(len(app.contexts)),
		pConn:       pConn,
		reference:   reference,
		dataOffset:  uint(refLength),
		contentList: []*content{},
	}

	return &pContext.identifier, nil
}

// Content returns the content by context
func (app *application) Content(context uint) (references.Content, error) {
	if pContext, ok := app.contexts[context]; ok {
		return pContext.reference.Content(), nil
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot return the Content instance", context)
	return nil, errors.New(str)
}

// Commits returns the commits on a context
func (app *application) Commits(context uint) (references.Commits, error) {
	if pContext, ok := app.contexts[context]; ok {
		if !pContext.reference.HasCommits() {
			str := fmt.Sprintf("the given context (%d) contains a reference that contains %d Commit instance", context, 0)
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
	content, err := app.Content(context)
	if err != nil {
		return nil, err
	}

	if !content.HasActive() {
		str := fmt.Sprintf("the ContentKey (hash: %s) could not be retrieved because the reference contains no active ContentKeys instance", hash.String())
		return nil, errors.New(str)
	}

	return content.Active().Fetch(hash)
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
	if pContext, ok := app.contexts[context]; ok {
		// find the latest commit:
		builder := app.commitBuilder.Create()
		if pContext.reference.HasCommits() {
			refCommit := pContext.reference.Commits().Latest()
			latestCommit, err := app.retrieveCommitByCommitReference(context, refCommit)
			if err != nil {
				return err
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
			return err
		}

		createdOn := time.Now().UTC()
		commit, err := builder.WithValues(values).CreatedOn(createdOn).Now()
		if err != nil {
			return err
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
			return err
		}

		commitBytes, err := app.commitContentAdapter.ToContent(commitContent)
		if err != nil {
			return err
		}

		// save the commit content on file:
		pointer, err := app.saveDataOnDisk(commitBytes, pContext.reference, pContext.pConn)
		if err != nil {
			return err
		}

		// build the commit reference:
		refCommit, err := app.referenceCommitBuilder.Create().WithHash(commitHash).WithPointer(pointer).Now()
		if err != nil {
			return err
		}

		// save the pointers in the commit references:
		commitsList := []references.Commit{}
		if pContext.reference.HasCommits() {
			commitsList = pContext.reference.Commits().List()
		}

		commitsList = append(commitsList, refCommit)
		commits, err := app.referenceCommitsBuilder.Create().WithList(commitsList).Now()
		if err != nil {
			return err
		}

		// get the pending content list:
		pendingList := []references.ContentKey{}
		content := pContext.reference.Content()
		if content.HasPendings() {
			pendingList = content.Pendings().List()
		}

		// save all content:
		for _, oneContent := range pContext.contentList {
			// save the content on file:
			pointer, err := app.saveDataOnDisk(oneContent.data, pContext.reference, pContext.pConn)
			if err != nil {
				return err
			}

			// build the content key:
			contentKey, err := app.referenceContentKeyBuilder.Create().WithHash(oneContent.hash).WithKind(oneContent.kind).WithContent(pointer).WithCommit(commitHash).Now()
			if err != nil {
				return err
			}

			//save the content key to the pending:
			pendingList = append(pendingList, contentKey)
		}

		pendingContentKeys, err := app.referenceContentKeysBuilder.Create().WithList(pendingList).Now()
		if err != nil {
			return err
		}

		updatedContentBuilder := app.referenceContentBuilder.Create().WithPendings(pendingContentKeys)
		if content.HasActive() {
			active := content.Active()
			updatedContentBuilder.WithActive(active)
		}

		if content.HasDeleted() {
			deleted := content.Deleted()
			updatedContentBuilder.WithDeleted(deleted)
		}

		updatedContent, err := updatedContentBuilder.Now()
		if err != nil {
			return err
		}

		updatedReference, err := app.referenceBuilder.Create().WithContent(updatedContent).WithCommits(commits).Now()
		if err != nil {
			return err
		}

		// update database on disk:
		return app.updateDatabaseOnDisk(pContext, updatedReference)
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be comitted", context)
	return errors.New(str)
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
	// rename the old database to a backup name:

	// rename the temporary database to the database name:

	// delete the backup database:
	return nil
}

func (app *application) saveDataOnDisk(data []byte, reference references.Reference, pConn *os.File) (references.Pointer, error) {
	// find the next pointer for the commit:
	from := reference.Next()

	// seek the file at the from byte:
	offset, err := pConn.Seek(from, 0)
	if err != nil {
		return nil, err
	}

	if offset != from {
		str := fmt.Sprintf("the offset was expected to be %d, %d returned after file seek", from, offset)
		return nil, errors.New(str)
	}

	// write the data on disk:
	amountWritten, err := pConn.Write(data)
	if err != nil {
		return nil, err
	}

	amountExpected := len(data)
	if amountExpected != amountWritten {
		str := fmt.Sprintf("%d bytes were expected to be written, %d actually written", amountExpected, amountWritten)
		return nil, errors.New(str)
	}

	// build the pointer:
	pointer, err := app.referencePointerBuilder.Create().From(uint(from)).WithLength(uint(amountWritten)).Now()
	if err != nil {
		return nil, err
	}

	return pointer, nil
}

// Push retrieves the commits from peers and chain then main our commit properlyusing the given configuration
func (app *application) Push(context uint, config configs.Config) error {
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
