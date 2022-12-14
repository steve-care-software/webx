package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/blockchains"
	"github.com/steve-care-software/webx/blockchains/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/blockchains/domain/blockchains/transactions"
	"github.com/steve-care-software/webx/blockchains/domain/contents/references"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

const (
	// PendingContentFlag represents the pending content flag
	PendingContentFlag uint8 = iota

	// ActiveContentFlag represents the active content flag
	ActiveContentFlag

	// DeletedContentFlag represents the deleted content flag
	DeletedContentFlag
)

const (
	// ChainBlockchainFlag represents the chain blockchain flag
	ChainBlockchainFlag uint8 = iota

	// BlockBlockchainFlag represents the block blockchain flag
	BlockBlockchainFlag

	// TransactionBlockchainFlag represents the transaction blockchain flag
	TransactionBlockchainFlag
)

// Application represents the read application
type Application interface {
	Database
	Reference
	Blockchain
	Content
}

// Blockchain represents the blockchain application
type Blockchain interface {
	Blockchain(context uint) (blockchains.Blockchain, error)
	BlockByHeight(context uint, height uint) (blocks.Block, error)
	BlockByHash(context uint, block hash.Hash) (blocks.Block, error)
	Pendings(context uint) (transactions.Transactions, error)
	Transactions(context uint, block hash.Hash) (transactions.Transactions, error)
	Transaction(context uint, block hash.Hash, trx hash.Hash) (transactions.Transaction, error)
	ReplaceTransaction(context uint, block hash.Hash, trx transactions.Transaction) error
}

// Reference represents the reference application
type Reference interface {
	ContentKeys(context uint, kind uint8) (references.ContentKeys, error)
	ContentKey(context uint, hash hash.Hash, flag uint8) (references.ContentKey, error)
	ContentKeyByTransaction(context uint, trx hash.Hash, flag uint8) (references.ContentKey, error)
	BlockchainKey(context uint, hash hash.Hash, flag uint8) (references.BlockchainKey, error)
}

// Database represents the database application
type Database interface {
	Delete(name string) error
	Open(name string, height int) (*uint, error)
	Cancel(context uint) error
	Commit(context uint) error
	Push(context uint) error
	Close(context uint) error
}

// Content represents the content application
type Content interface {
	Read(context uint, pointer references.Pointer) ([]byte, error)
	ReadByHash(content uint, hash hash.Hash) ([]byte, error)
	ReadAll(context uint, pointers []references.Pointer) ([][]byte, error)
	ReadAllByHashes(context uint, hashes []hash.Hash) ([][]byte, error)
	Write(context uint, hash hash.Hash, data []byte, kind uint8) error
}
