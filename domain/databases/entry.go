package databases

import (
	"github.com/steve-care-software/webx/domain/databases/blockchains/transactions"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type entry struct {
	entity  entities.Entity
	trx     transactions.Transaction
	pointer Pointer
	content []byte
	kind    uint8
}

func createEntry(
	entity entities.Entity,
	trx transactions.Transaction,
	pointer Pointer,
	content []byte,
	kind uint8,
) Entry {
	out := entry{
		entity:  entity,
		trx:     trx,
		pointer: pointer,
		content: content,
		kind:    kind,
	}

	return &out
}

// Entity returns the entity
func (obj *entry) Entity() entities.Entity {
	return obj.entity
}

// Transaction returns the trx
func (obj *entry) Transaction() transactions.Transaction {
	return obj.trx
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
