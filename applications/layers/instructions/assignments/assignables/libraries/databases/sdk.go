package databases

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases"
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
