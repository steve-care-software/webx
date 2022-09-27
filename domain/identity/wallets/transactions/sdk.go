package transactions

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets/units"
)

// Builder represents a transactions builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents a transactions
type Transactions interface {
	List() []Transaction
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithID(id uuid.UUID) TransactionBuilder
	WithUnit(unit units.Unit) TransactionBuilder
	CreatedOn(createdOn time.Time) TransactionBuilder
	WithDetails(details string) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	ID() uuid.UUID
	Unit() units.Unit
	CreatedOn() time.Time
	HasDetails() bool
	Details() string
}
