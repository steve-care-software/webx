package blockchains

import (
	applications "github.com/steve-care-software/webx/roots/applications/blockchains"
	"github.com/steve-care-software/webx/roots/domain/blockchains/blockchains"
	"github.com/steve-care-software/webx/roots/domain/blockchains/blockchains/blocks"
	"github.com/steve-care-software/webx/roots/domain/blockchains/blockchains/transactions"
	"github.com/steve-care-software/webx/roots/domain/blockchains/contents/references"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type application struct {
	dirPath string
}

func createApplication(
	dirPath string,
) applications.Application {
	out := application{
		dirPath: dirPath,
	}

	return &out
}

// Delete deletes an existing database
func (app *application) Delete(name string) error {
	return nil
}

// Open opens a context at height, height is -1 if the head is requested
func (app *application) Open(name string, height int) (*uint, error) {
	return nil, nil
}

// ContentKeys returns the contentKey kind on a context
func (app *application) ContentKeys(context uint, kind uint8) (references.ContentKeys, error) {
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

// ListByKind returns the list by kind
func (app *application) ListByKind(context uint, kind uint, index uint, amount uint) ([]hash.Hash, uint, error) {
	return nil, 0, nil
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
func (app *application) Write(context uint, hash hash.Hash, data []byte, kind uint8) error {
	return nil
}

// Cancel cancels a context
func (app *application) Cancel(context uint) error {
	return nil
}

// Commit commits a context
func (app *application) Commit(context uint) error {
	return nil
}

// Push pushes a context
func (app *application) Push(context uint) error {
	return nil
}

// Close closes a context
func (app *application) Close(context uint) error {
	return nil
}
