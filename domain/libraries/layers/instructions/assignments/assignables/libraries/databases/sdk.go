package databases

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/databases/actions"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/databases/transactions"
)

// Database represents a database instruction
type Database interface {
	IsSkeleton() bool
	IsBegin() bool
	IsList() bool
	List() string
	IsRetrieve() bool
	Retrieve() Retrieve
	IsCommit() bool
	Commit() string
	IsAction() bool
	Action() actions.Action
	IsTransaction() bool
	Transaction() transactions.Transaction
}

// Retrieve represents a retrieve
type Retrieve interface {
	Path() string
	Hash() string
}
