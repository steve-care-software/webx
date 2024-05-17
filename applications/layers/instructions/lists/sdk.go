package lists

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, list lists.List) (stacks.Assignment, *uint, error)
}
