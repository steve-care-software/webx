package inserts

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable inserts.Insert) ([]stacks.Assignable, *uint, error)
}
