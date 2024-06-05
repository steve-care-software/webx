package modifications

import (
	databases_modifications "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	modificationBuilder := databases_modifications.NewModificationBuilder()
	return createApplication(
		assignableBuilder,
		modificationBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable modifications.Modification) (stacks.Assignable, *uint, error)
}
