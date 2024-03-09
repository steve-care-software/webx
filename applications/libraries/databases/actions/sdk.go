package actions

import "github.com/steve-care-software/datastencil/domain/orms"

// Application represents the action application
type Application interface {
	Amount() (uint, error)
	Retrieve(index uint) (orms.Actions, error)
	Rollback(toIndex uint) error
}
