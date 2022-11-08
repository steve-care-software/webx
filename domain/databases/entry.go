package databases

import (
	"github.com/steve-care-software/webx/domain/databases/blockchains/transactions"
)

type entry struct {
	content []byte
	kind    uint8
	trx     transactions.Transaction
}

func createEntry(
	content []byte,
	kind uint8,
) Entry {
	return createEntryInternally(content, kind, nil)
}

func createEntryWithTransaction(
	content []byte,
	kind uint8,
	trx transactions.Transaction,
) Entry {
	return createEntryInternally(content, kind, trx)
}

func createEntryInternally(
	content []byte,
	kind uint8,
	trx transactions.Transaction,
) Entry {
	out := entry{
		content: content,
		kind:    kind,
		trx:     trx,
	}

	return &out
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
