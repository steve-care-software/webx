package accounts

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/accounts/updates"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the account adapter
type Adapter interface {
	ToBytes(ins Account) ([]byte, error)
	ToInstance(bytes []byte) (Account, error)
}

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithInsert(insert inserts.Insert) Builder
	WithUpdate(update updates.Update) Builder
	WithDelete(delete string) Builder
	Now() (Account, error)
}

// Account represents an account instruction
type Account interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() inserts.Insert
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() string
}
