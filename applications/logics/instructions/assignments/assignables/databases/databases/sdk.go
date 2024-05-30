package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable databases.Database) (stacks.Assignable, *uint, error)
}
