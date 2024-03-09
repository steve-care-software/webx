package accounts

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/deletes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/updates"
)

// Builder represents an account builder
type Builder interface {
	Create() Builder
	WithUpate(update updates.Update) Builder
	WithDelete(delete deletes.Delete) Builder
	Now() (Account, error)
}

// Account represents an account instruction
type Account interface {
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() deletes.Delete
}
