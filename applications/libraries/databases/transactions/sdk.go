package transactions

import "github.com/steve-care-software/datastencil/domain/orms"

// Application represents a transaction application
type Application interface {
	Actions() (orms.Actions, error)
	Insert(resource orms.Resource) error
	Delete(resource orms.Pointer) error
}
