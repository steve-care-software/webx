package deletes

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignment deletes.Delete) (stacks.Assignment, *uint, error)
}
