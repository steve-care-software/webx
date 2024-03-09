package cryptography

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable cryptography.Cryptography) (stacks.Assignable, error)
}
