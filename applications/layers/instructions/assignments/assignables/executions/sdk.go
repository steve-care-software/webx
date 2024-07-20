package executions

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/executions/retrieves"
	instruction_execution "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execExecuteApp executes.Application,
	execInitApp inits.Application,
	execRetrieveApp retrieves.Application,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	assignablesBuilder := stacks.NewAssignablesBuilder()
	return createApplication(
		execExecuteApp,
		execInitApp,
		execRetrieveApp,
		assignableBuilder,
		assignablesBuilder,
	)
}

// Application represents an execution application
type Application interface {
	Execute(frame stacks.Frame, assignable instruction_execution.Execution) (stacks.Assignable, *uint, error)
}
