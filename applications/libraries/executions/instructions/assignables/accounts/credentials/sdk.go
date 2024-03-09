package credentials

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable credentials.Credentials) (stacks.Assignable, error)
}
