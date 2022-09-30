package transactions

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/composers"
)

// Transactions represents transactions
type Transactions interface {
	List() []Transaction
}

// Transaction represents a database transaction
type Transaction interface {
	Criteria() criterias.Criteria
	HasComposer() bool
	Composer() composers.Composer
}
