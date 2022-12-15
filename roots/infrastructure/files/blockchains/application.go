package blockchains

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	applications "github.com/steve-care-software/webx/roots/applications/blockchains"
	"github.com/steve-care-software/webx/roots/domain/blockchains/blockchains"
	"github.com/steve-care-software/webx/roots/domain/blockchains/blockchains/blocks"
	"github.com/steve-care-software/webx/roots/domain/blockchains/blockchains/transactions"
	"github.com/steve-care-software/webx/roots/domain/blockchains/contents/references"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type application struct {
	referenceAdapter   references.Adapter
	contentKeyBuilder  references.ContentKeyBuilder
	transactionBuilder transactions.TransactionBuilder
	dirPath            string
	contexts           map[uint]*context
}

func createApplication(
	referenceAdapter references.Adapter,
	contentKeyBuilder references.ContentKeyBuilder,
	transactionBuilder transactions.TransactionBuilder,
	dirPath string,
) applications.Application {
	out := application{
		referenceAdapter:   referenceAdapter,
		contentKeyBuilder:  contentKeyBuilder,
		transactionBuilder: transactionBuilder,
		dirPath:            dirPath,
		contexts:           map[uint]*context{},
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

// ContentKey returns the contentKey by hash and flag on a context
func (app *application) ContentKey(context uint, hash hash.Hash, flag uint8) (references.ContentKey, error) {
	return nil, nil
}

// ContentKeyByTransaction returns the contentKey by transaction hash and flag on a context
func (app *application) ContentKeyByTransaction(context uint, trx hash.Hash, flag uint8) (references.ContentKey, error) {
	return nil, nil
}

// BlockchainKey returns the blockchainKey by hash and flag on a context
func (app *application) BlockchainKey(context uint, hash hash.Hash, flag uint8) (references.BlockchainKey, error) {
	return nil, nil
}

// Blockchain returns the blockchain on a context
func (app *application) Blockchain(context uint) (blockchains.Blockchain, error) {
	return nil, nil
}

// BlockByHeight returns the block by height
func (app *application) BlockByHeight(context uint, height uint) (blocks.Block, error) {
	return nil, nil
}

// BlockByHash returns the block by hash
func (app *application) BlockByHash(context uint, block hash.Hash) (blocks.Block, error) {
	return nil, nil
}

// Pendings returns the pending transactions
func (app *application) Pendings(context uint) (transactions.Transactions, error) {
	return nil, nil
}

// Transactions returns the transactions by block
func (app *application) Transactions(context uint, block hash.Hash) (transactions.Transactions, error) {
	return nil, nil
}

// Transaction returns the transaction by block and trx hash
func (app *application) Transaction(context uint, block hash.Hash, trx hash.Hash) (transactions.Transaction, error) {
	return nil, nil
}

// ReplaceTransaction replaces a transaction
func (app *application) ReplaceTransaction(context uint, block hash.Hash, trx transactions.Transaction) error {
	return nil
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
	if pContext, ok := app.contexts[context]; ok {
		for _, oneContent := range pContext.contentList {

			// create a new transaction:

			// find the latest block:

			// save the transaction:

			// find the next point for the data kind:

			// build a new content key:
			//contentKey, err := app.contentKeyBuilder.Create().WithHash(oneContent.hash).WithKind(pContext.kind)

			// save the content key to the reference:

			// save the reference on disk:
		}
	}

	str := fmt.Sprintf("the given context (%d) does not exists and therefore cannot be comitted", context)
	return errors.New(str)
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
