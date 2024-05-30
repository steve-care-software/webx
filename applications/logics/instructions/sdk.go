package instructions

import (
	application_assignments "github.com/steve-care-software/datastencil/applications/logics/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/interruptions"
	results_failures "github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results/interruptions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execAssignmentApp application_assignments.Application,
) Application {
	stackBuilder := stacks.NewBuilder()
	framesBuilder := stacks.NewFramesBuilder()
	frameBuilder := stacks.NewFrameBuilder()
	assignmentsBuilder := stacks.NewAssignmentsBuilder()
	interruptionBuilder := interruptions.NewBuilder()
	failureBuilder := results_failures.NewBuilder()
	return createApplication(
		execAssignmentApp,
		nil,
		nil,
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
