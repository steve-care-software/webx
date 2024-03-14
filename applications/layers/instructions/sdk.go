package instructions

import (
	application_accounts "github.com/steve-care-software/datastencil/applications/layers/instructions/accounts"
	application_assignments "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments"
	application_databases "github.com/steve-care-software/datastencil/applications/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions"
	results_failures "github.com/steve-care-software/datastencil/domain/instances/commands/results/interruptions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execAccountApp application_accounts.Application,
	execAssignmentApp application_assignments.Application,
	execDatabaseApp application_databases.Application,
) Application {
	stackBuilder := stacks.NewBuilder()
	framesBuilder := stacks.NewFramesBuilder()
	frameBuilder := stacks.NewFrameBuilder()
	assignmentsBuilder := stacks.NewAssignmentsBuilder()
	interruptionBuilder := interruptions.NewBuilder()
	failureBuilder := results_failures.NewBuilder()
	return createApplication(
		execAccountApp,
		execAssignmentApp,
		execDatabaseApp,
		stackBuilder,
		framesBuilder,
		frameBuilder,
		assignmentsBuilder,
		interruptionBuilder,
		failureBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(stack stacks.Stack, instructions instructions.Instructions) (stacks.Stack, interruptions.Interruption, error)
}
