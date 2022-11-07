package databases

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/databases/blockchains/transactions"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type entryBuilder struct {
	entity  entities.Entity
	trx     transactions.Transaction
	pointer Pointer
	content []byte
	pKind   *uint8
}

func createEntryBuilder() EntryBuilder {
	out := entryBuilder{
		entity:  nil,
		trx:     nil,
		pointer: nil,
		content: nil,
		pKind:   nil,
	}

	return &out
}

// Create initializes the bulder
func (app *entryBuilder) Create() EntryBuilder {
	return createEntryBuilder()
}

// WithEntity adds an entity to the builder
func (app *entryBuilder) WithEntity(entity entities.Entity) EntryBuilder {
	app.entity = entity
	return app
}

// WithTransaction adds a transaction to the builder
func (app *entryBuilder) WithTransaction(transaction transactions.Transaction) EntryBuilder {
	app.trx = transaction
	return app
}

// WithPointer adds a pointer to the builder
func (app *entryBuilder) WithPointer(pointer Pointer) EntryBuilder {
	app.pointer = pointer
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
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Entry instance")
	}

	if app.trx == nil {
		return nil, errors.New("the transaction is mandatory in order to build an Entry instance")
	}

	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build an Entry instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Entry instance")
	}

	if app.pKind == nil {
		return nil, errors.New("the kind is mandatory in order to build an Entry instance")
	}

	if *app.pKind <= KindEntry {
		str := fmt.Sprintf("the kind flag must be a uint8 with a value between %d and %d, %d provided", KindBlockchain, KindEntry, *app.pKind)
		return nil, errors.New(str)
	}

	return createEntry(app.entity, app.trx, app.pointer, app.content, *app.pKind), nil

}
