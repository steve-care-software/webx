package constants

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable constants.Constant) (stacks.Assignable, error)
}
