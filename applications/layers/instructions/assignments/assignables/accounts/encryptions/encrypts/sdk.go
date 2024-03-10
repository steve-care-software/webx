package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, error)
}
