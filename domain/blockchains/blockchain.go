package blockchains

import (
	"time"

	"github.com/steve-care-software/webx/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/domain/blockchains/transactions"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type blockchain struct {
	reference hash.Hash
	head      blocks.Block
	createdOn time.Time
	pendings  transactions.Transactions
}

func createBlockchain(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
) Blockchain {
	return createBlockchainInternally(reference, head, createdOn, nil)
}

func createBlockchainWithPendings(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
	pendings transactions.Transactions,
) Blockchain {
	return createBlockchainInternally(reference, head, createdOn, pendings)
}

func createBlockchainWithPendingsAndConnections(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
	pendings transactions.Transactions,
) Blockchain {
	return createBlockchainInternally(reference, head, createdOn, pendings)
}

func createBlockchainInternally(
	reference hash.Hash,
	head blocks.Block,
	createdOn time.Time,
	pendings transactions.Transactions,
) Blockchain {
	out := blockchain{
		reference: reference,
		head:      head,
		createdOn: createdOn,
		pendings:  pendings,
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
