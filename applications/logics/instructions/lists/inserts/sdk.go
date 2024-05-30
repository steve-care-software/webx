package inserts

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable inserts.Insert) (stacks.Assignment, *uint, error)
}
