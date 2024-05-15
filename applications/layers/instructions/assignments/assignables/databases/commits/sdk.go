package commits

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable commits.Commit) (stacks.Assignable, *uint, error)
}
