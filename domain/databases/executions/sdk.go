package executions

import (
	"time"

	"github.com/steve-care-software/syntax/domain/databases/transactions"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	Previous() hash.Hash
	Content() Content
	IsSuccessful() bool
	CreatedOn() time.Time
}

// Content represents an execution content
type Content interface {
	IsSingle() bool
	Single() transactions.Transaction
	IsAtomic() bool
	Atomic() transactions.Transactions
}
