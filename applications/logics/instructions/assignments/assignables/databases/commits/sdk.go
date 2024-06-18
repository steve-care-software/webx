package commits

import (
	databases_commits "github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	actionsBuilder := actions.NewBuilder()
	commitBuilder := databases_commits.NewBuilder()
	return createApplication(
		assignableBuilder,
		actionsBuilder,
		commitBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable commits.Commit) (stacks.Assignable, *uint, error)
}
