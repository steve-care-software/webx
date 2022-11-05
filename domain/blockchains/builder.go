package blockchains

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/domain/blockchains/transactions"
)

type builder struct {
	pReference *hash.Hash
	head       blocks.Block
	pCreatedOn *time.Time
	pendings   transactions.Transactions
}

func createBuilder() Builder {
	out := builder{
		pReference: nil,
		head:       nil,
		pCreatedOn: nil,
		pendings:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithReference adds a reference to the builder
func (app *builder) WithReference(reference hash.Hash) Builder {
	app.pReference = &reference
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head blocks.Block) Builder {
	app.head = head
	return app
}

// WithPendings add pending transactions to the builder
func (app *builder) WithPendings(pendings transactions.Transactions) Builder {
	app.pendings = pendings
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Blockchain instance
func (app *builder) Now() (Blockchain, error) {
	if app.pReference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Blockchain instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Blockchain instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Blockchain instance")
	}

	if app.pendings != nil {
		return createBlockchainWithPendings(*app.pReference, app.head, *app.pCreatedOn, app.pendings), nil
	}

	return createBlockchain(*app.pReference, app.head, *app.pCreatedOn), nil
}
