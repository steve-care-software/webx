package databases

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/repositories"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/services"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execRepositoryApp repositories.Application,
	execServiceApp services.Application,
) Application {
	return createApplication(
		execRepositoryApp,
		execServiceApp,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable databases.Database) (stacks.Assignable, *uint, error)
}
