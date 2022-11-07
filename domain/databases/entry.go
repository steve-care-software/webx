package databases

import (
	"github.com/steve-care-software/webx/domain/databases/blockchains/transactions"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type entry struct {
	entity  entities.Entity
	pointer Pointer
	content []byte
	kind    uint8
	trx     transactions.Transaction
}

func createEntry(
	entity entities.Entity,
	pointer Pointer,
	content []byte,
	kind uint8,
) Entry {
	return createEntryInternally(entity, pointer, content, kind, nil)
}

func createEntryWithTransaction(
	entity entities.Entity,
	pointer Pointer,
	content []byte,
	kind uint8,
	trx transactions.Transaction,
) Entry {
	return createEntryInternally(entity, pointer, content, kind, trx)
}

func createEntryInternally(
	entity entities.Entity,
	pointer Pointer,
	content []byte,
	kind uint8,
	trx transactions.Transaction,
) Entry {
	out := entry{
		entity:  entity,
		pointer: pointer,
		content: content,
		kind:    kind,
		trx:     trx,
	}

	return &out
}

// Entity returns the entity
func (obj *entry) Entity() entities.Entity {
	return obj.entity
}

// Pointer returns the pointer
func (obj *entry) Pointer() Pointer {
	return obj.pointer
}

// Content returns the content
func (obj *entry) Content() []byte {
	return obj.content
}

// Kind returns the kind
func (obj *entry) Kind() uint8 {
	return obj.kind
}

// HasTransaction returns true if there is a trx, false otherwise
func (obj *entry) HasTransaction() bool {
	return obj.trx != nil
}

// Transaction returns the trx
func (obj *entry) Transaction() transactions.Transaction {
	return obj.trx
}
