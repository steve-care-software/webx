package actions

import "github.com/steve-care-software/datastencil/domain/commits/actions"

// Application represents the action application
type Application interface {
	Amount() (uint, error)
	Retrieve(index uint) (actions.Actions, error)
	Rollback(toIndex uint) error
}
