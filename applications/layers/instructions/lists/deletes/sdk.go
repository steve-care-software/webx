package deletes

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignment deletes.Delete) ([]stacks.Assignable, *uint, error)
}
