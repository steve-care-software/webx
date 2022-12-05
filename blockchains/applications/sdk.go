package applications

import (
	"github.com/steve-care-software/webx/blockchains/domain/blockchains"
	"github.com/steve-care-software/webx/blockchains/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/blockchains/domain/blockchains/transactions"
	"github.com/steve-care-software/webx/blockchains/domain/contents/references"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// Application represents the read application
type Application interface {
	New(name string) error
	Database
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

// Database represents the database application
type Database interface {
	Delete(name string) error
	List() ([]string, error)
	Open(name string, height int) (*uint, error)
	Cancel(context uint) error
	Commit(context uint) error
	Push(context uint) error
	Close(context uint) error
}

// Content represents the content application
type Content interface {
	Read(context uint, pointer references.Pointer) ([]byte, error)
	ReadAll(context uint, pointers []references.Pointer) ([][]byte, error)
	Write(data []byte) error
	WriteAll(data [][]byte) error
}
