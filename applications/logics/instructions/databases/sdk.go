package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	database_instruction "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	databaseService databases.Service,
) Application {
	return createApplication(
		databaseService,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignment database_instruction.Database) (*uint, error)
}
