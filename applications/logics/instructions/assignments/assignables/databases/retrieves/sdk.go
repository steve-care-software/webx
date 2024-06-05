package retrieves

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	repository databases.Repository,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	assignablesBuilder := stacks.NewAssignablesBuilder()
	return createApplication(
		repository,
		assignableBuilder,
		assignablesBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable retrieves.Retrieve) (stacks.Assignable, *uint, error)
}
