package transactions

import (
	"github.com/steve-care-software/datastencil/domain/commits/actions"
	"github.com/steve-care-software/datastencil/domain/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources"
)

// Application represents a transaction application
type Application interface {
	Actions() (actions.Actions, error)
	Insert(resource resources.Resource) error
	Delete(resource pointers.Pointer) error
}
