package accounts

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/updates"
)

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithInsert(isnert inserts.Insert) Builder
	WithUpdate(update updates.Update) Builder
	WithDelete(delete string) Builder
	Now() (Account, error)
}

// Account represents an account instruction
type Account interface {
	IsInsert() bool
	Insert() inserts.Insert
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() string
}
