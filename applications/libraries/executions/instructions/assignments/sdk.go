package assignments

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignment assignments.Assignment) (stacks.Assignment, error)
}
