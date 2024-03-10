package libraries

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable libraries.Library) (stacks.Assignable, error)
}
