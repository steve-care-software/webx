package deletes

import (
	databases_deletes "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	deleteBuilder := databases_deletes.NewBuilder()
	return createApplication(
		assignableBuilder,
		deleteBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable deletes.Delete) (stacks.Assignable, *uint, error)
}
