package files

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/commits"
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
	referenceContentBuilder     references.ContentBuilder
	referenceContentKeysBuilder references.ContentKeysBuilder
	referenceContentKeyBuilder  references.ContentKeyBuilder
	referenceCommitsBuilder     references.CommitsBuilder
	referenceCommitBuilder      references.CommitBuilder
	referencePointerBuilder     references.PointerBuilder
	hashTreeBuilder             hashtrees.Builder
	dirPath                     string
	contexts                    map[uint]*context
}

func createApplication(
	commitBuilder commits.Builder,
	commitContentAdapter commit_contents.Adapter,
	commitContentBuilder commit_contents.Builder,
	referenceAdapter references.Adapter,
	referenceBuilder references.Builder,
	referenceContentBuilder references.ContentBuilder,
	referenceContentKeysBuilder references.ContentKeysBuilder,
	referenceContentKeyBuilder references.ContentKeyBuilder,
	referenceCommitsBuilder references.CommitsBuilder,
	referenceCommitBuilder references.CommitBuilder,
	referencePointerBuilder references.PointerBuilder,
	hashTreeBuilder hashtrees.Builder,
	dirPath string,
) applications.Application {
	out := application{
		commitBuilder:               commitBuilder,
		commitContentAdapter:        commitContentAdapter,
		commitContentBuilder:        commitContentBuilder,
		referenceAdapter:            referenceAdapter,
		referenceBuilder:            referenceBuilder,
		referenceContentBuilder:     referenceContentBuilder,
		referenceContentKeysBuilder: referenceContentKeysBuilder,
		referenceContentKeyBuilder:  referenceContentKeyBuilder,
		referenceCommitsBuilder:     referenceCommitsBuilder,
		referenceCommitBuilder:      referenceCommitBuilder,
		referencePointerBuilder:     referencePointerBuilder,
		hashTreeBuilder:             hashTreeBuilder,
		dirPath:                     dirPath,
		contexts:                    map[uint]*context{},
	}

	return &out
}

// Delete deletes an existing database
func (app *application) Delete(name string) error {
	path := filepath.Join(app.dirPath, name)
	pInfo, err := os.Stat(name)
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
	refContentBytes := make([]byte, refLength)
	refContentAmount, err := pConn.ReadAt(refContentBytes, int64(refLength))
	if err != nil {
		return nil, err
	}

	if refContentAmount != int(refLength) {
		str := fmt.Sprintf("%d bytes were expected to be read when reading the reference bytes, %d actually read", refLength, refContentAmount)
		return nil, errors.New(str)
	}

	// convert the content to a reference instance:
	reference, err := app.referenceAdapter.ToReference(refContentBytes)
	if err != nil {
		return nil, err
	}

	pContext := &context{
		identifier:  uint(len(app.contexts)),
		pConn:       pConn,
		reference:   reference,
		contentList: []*content{},
	}

	return &pContext.identifier, nil
}

// ContentKeys returns the contentKey kind on a context
func (app *application) ContentKeys(context uint, kind uint) (references.ContentKeys, error) {
	return nil, nil
}

// ContentKeysByCommit returns the contentKeys by commit on a context
func (app *application) ContentKeysByCommit(context uint, commit hash.Hash) (references.ContentKeys, error) {
	return nil, nil
}

// ContentKey returns the contentKey by hash and flag on a context
func (app *application) ContentKey(context uint, hash hash.Hash, flag uint8) (references.ContentKey, error) {
	return nil, nil
}

// Commits returns the commits on a context
func (app *application) Commits(context uint) (references.Commits, error) {
	return nil, nil
}

// Latest retrieves the latest commit
func (app *application) Latest() (commits.Commit, error) {
	return nil, nil
}

// Retrieve retrieves the commit by hash
func (app *application) Retrieve(hash hash.Hash) (commits.Commit, error) {
	return nil, nil
}

// Read reads a pointer on a context
func (app *application) Read(context uint, pointer references.Pointer) ([]byte, error) {
	return nil, nil
}

// ReadByHash reads content by hash
func (app *application) ReadByHash(content uint, hash hash.Hash) ([]byte, error) {
	return nil, nil
}

// ReadAll read pointers on a context
func (app *application) ReadAll(context uint, pointers []references.Pointer) ([][]byte, error) {
	return nil, nil
}

// ReadAllByHashes reads content by hashes
func (app *application) ReadAllByHashes(context uint, hashes []hash.Hash) ([][]byte, error) {
	return nil, nil
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
	// find the latest commit:
	latestCommit, _ := app.Latest()
	builder := app.commitBuilder.Create()
	if latestCommit != nil {
		builder.WithParent(latestCommit)
	}

	blocks := [][]byte{}
	if pContext, ok := app.contexts[context]; ok {
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
		builder := app.commitContentBuilder.Create().WithHash(commitHash).WithValues(commitValues).CreatedOn(createdOn)
		if commit.HasParent() {
			commitParent := commit.Parent().Hash()
			builder.WithParent(commitParent)
		}

		commitContent, err := builder.Now()
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

// Push pushes a context
func (app *application) Push(context uint) error {
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
