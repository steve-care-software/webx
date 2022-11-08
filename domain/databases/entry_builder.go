package databases

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/databases/blockchains/transactions"
)

type entryBuilder struct {
	trx     transactions.Transaction
	content []byte
	pKind   *uint8
}

func createEntryBuilder() EntryBuilder {
	out := entryBuilder{
		trx:     nil,
		content: nil,
		pKind:   nil,
	}

	return &out
}

// Create initializes the bulder
func (app *entryBuilder) Create() EntryBuilder {
	return createEntryBuilder()
}

// WithTransaction adds a transaction to the builder
func (app *entryBuilder) WithTransaction(transaction transactions.Transaction) EntryBuilder {
	app.trx = transaction
	return app
}

// WithContent adds content to the builder
func (app *entryBuilder) WithContent(content []byte) EntryBuilder {
	app.content = content
	return app
}

// WithKind adds a kind to the builder
func (app *entryBuilder) WithKind(kind uint8) EntryBuilder {
	app.pKind = &kind
	return app
}

// Now builds a new Entry instance
func (app *entryBuilder) Now() (Entry, error) {
	if app.trx == nil {
		return nil, errors.New("the transaction is mandatory in order to build an Entry instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Entry instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build an Entry instance")
	}

	kind := *app.pKind
	if kind <= KindEntry {
		str := fmt.Sprintf("the kind must be a uint8 with a value between %d and %d, %d provided", KindBlockchain, KindEntry, kind)
		return nil, errors.New(str)
	}

	if app.trx != nil {
		return createEntryWithTransaction(app.content, *app.pKind, app.trx), nil
	}

	if kind != KindBlockchain && kind != KindBlockchainBlock && kind != KindBlockchainTransaction {
		return nil, errors.New("the entry must have a transaction when the kind is NOT blockchain related")
	}

	return createEntry(app.content, *app.pKind), nil

}
