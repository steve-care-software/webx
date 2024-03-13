package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"

	application_accounts "github.com/steve-care-software/datastencil/applications/layers/instructions/accounts"
	application_assignments "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments"
	application_databases "github.com/steve-care-software/datastencil/applications/layers/instructions/databases"
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
	failureBuilder := results.NewFailureBuilder()
	return createApplication(
		execAccountApp,
		execAssignmentApp,
		execDatabaseApp,
		stackBuilder,
		framesBuilder,
		frameBuilder,
		assignmentsBuilder,
		failureBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(stack stacks.Stack, instructions instructions.Instructions) (bool, stacks.Stack, results.Failure, error)
}
