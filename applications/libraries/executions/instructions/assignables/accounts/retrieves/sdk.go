package retrieves

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable retrieves.Retrieve) (stacks.Assignable, error)
}
