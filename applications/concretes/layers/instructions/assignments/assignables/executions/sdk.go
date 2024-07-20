package executions

import (
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/executions/retrieves"
	instruction_execution "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execAmountApp amounts.Application,
	execBeginApp begins.Application,
	execExecuteApp executes.Application,
	execHeadApp heads.Application,
	execInitApp inits.Application,
	execRetrieveApp retrieves.Application,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	assignablesBuilder := stacks.NewAssignablesBuilder()
	return createApplication(
		execAmountApp,
		execBeginApp,
		execExecuteApp,
		execHeadApp,
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
