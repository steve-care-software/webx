package blockchains

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/blockchains/blocks"
	"github.com/steve-care-software/webx/engine/domain/blockchains"
	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions"
	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions/contents"
	"github.com/steve-care-software/webx/engine/domain/blockchains/roots"
	"github.com/steve-care-software/webx/engine/domain/hash"
)

// Application represents the blockchain application
type Application interface {
	Begin() error
	List() ([]string, error)                                                               // lists the blockchain package names
	Retrieve(packageName string) (blockchains.Blockchain, error)                           // retrieve the blockchain
	Create(root roots.Root) (blockchains.Blockchain, error)                                // creates a blockchain
	Insert(blockchain blockchains.Blockchain) error                                        // inserts a bockchain
	Delete(packageName string) error                                                       // deletes a blockchain
	Fork(packageName string, newHead blocks.Block) error                                   // change the head of the bockchain of that package
	RetrieveBlock(blockchain blockchains.Blockchain, blockHash hash.Hash)                  // retrieve the block of a blockchain
	Transact(blockchain blockchains.Blockchain, content contents.Content, pk string) error // add a transaction to the provided blockchain
	Queue(blockchain blockchains.Blockchain) (transactions.Transactions, error)            // returns the transaction queue
	Mine(blockchain blockchains.Blockchain, pk string) (blocks.Block, error)               // mines a block and returns it
	InsertBlock(blockchain blockchains.Blockchain, block blocks.Block) error               // inserts a block to the database
	InsertPeer(blockchain blockchains.Blockchain, host string) error                       // inserts a peer to the blockchain
	DeletePeer(blockchain blockchains.Blockchain, host string) error                       // deletes a peer from the blockchain
	Commit() error                                                                         // commits the transaction to the blockchain database
	Cancel() error                                                                         // cancels the transaction
}
