package databases

import (
	application_deletes "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/deletes"
	application_inserts "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/inserts"
	application_reverts "github.com/steve-care-software/datastencil/applications/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execDeleteApp application_deletes.Application,
	execInsertApp application_inserts.Application,
	execRevertApp application_reverts.Application,
	service instances.Service,
) Application {
	return createApplication(
		execDeleteApp,
		execInsertApp,
		execRevertApp,
		service,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction databases.Database) (*uint, error)
}
