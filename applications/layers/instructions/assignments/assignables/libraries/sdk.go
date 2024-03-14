package libraries

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/libraries"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execCompilerApp compilers.Application,
	execDatabaseApp databases.Application,
) Application {
	return createApplication(
		execCompilerApp,
		execDatabaseApp,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable libraries.Library) (stacks.Assignable, *uint, error)
}
