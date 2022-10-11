package selects

import (
	"github.com/steve-care-software/syntax/domain/databases/transactions"
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
)

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Application, error)
}

// Application represents a selected database application
type Application interface {
	Select(criteria criterias.Criteria) ([]byte, error)
	AtomicTransact(trx transactions.Transactions) error
	Transact(trx transactions.Transaction) error
}
