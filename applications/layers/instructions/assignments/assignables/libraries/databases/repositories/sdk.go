package repositories

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/skeletons"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	repository instances.Repository,
	skeleton skeletons.Skeleton,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		repository,
		assignableBuilder,
		skeleton,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable repositories.Repository) (stacks.Assignable, *uint, error)
}
