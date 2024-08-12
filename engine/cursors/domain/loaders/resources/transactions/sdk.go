package transactions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
)

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	List() []Transaction
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() inserts.Insert
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() deletes.Delete
}
