package blockchains

import (
	"net/url"
	"time"

	"github.com/steve-care-software/syntax/domain/blockchains/blocks"
	"github.com/steve-care-software/syntax/domain/blockchains/transactions"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type blockchain struct {
	reference   hash.Hash
	head        blocks.Block
	createdOn   time.Time
	pendings    transactions.Transactions
	connections []*url.URL
}

func createBlockchain(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
) Blockchain {
	return createBlockchainInternally(reference, head, createdOn, nil, nil)
}

func createBlockchainWithPendings(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
	pendings transactions.Transactions,
) Blockchain {
	return createBlockchainInternally(reference, head, createdOn, pendings, nil)
}

func createBlockchainWithConnections(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
	connections []*url.URL,
) Blockchain {
	return createBlockchainInternally(reference, head, createdOn, nil, connections)
}

func createBlockchainWithPendingsAndConnections(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
	pendings transactions.Transactions,
	connections []*url.URL,
) Blockchain {
	return createBlockchainInternally(reference, head, createdOn, pendings, connections)
}

func createBlockchainInternally(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
	pendings transactions.Transactions,
	connections []*url.URL,
) Blockchain {
	out := blockchain{
		reference:   reference,
		head:        head,
		createdOn:   createdOn,
		pendings:    pendings,
		connections: connections,
	}

	return &out
}

// Reference returns the reference
func (obj *blockchain) Reference() hash.Hash {
	return obj.reference
}

// Head returns the head block
func (obj *blockchain) Head() blocks.Block {
	return obj.head
}

// CreatedOn returns the creation time
func (obj *blockchain) CreatedOn() time.Time {
	return obj.createdOn
}

// HasPendings returns true if there is pending transactions, false otherwise
func (obj *blockchain) HasPendings() bool {
	return obj.pendings != nil
}

// Pendings returns the pending transactions, if any
func (obj *blockchain) Pendings() transactions.Transactions {
	return obj.pendings
}

// HasConnections returns true if there is connections, false otherwise
func (obj *blockchain) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *blockchain) Connections() []*url.URL {
	return obj.connections
}
