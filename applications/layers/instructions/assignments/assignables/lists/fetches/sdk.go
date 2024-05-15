package fetches

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable fetches.Fetch) (stacks.Assignable, *uint, error)
}
