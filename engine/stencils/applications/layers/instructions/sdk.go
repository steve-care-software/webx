package instructions

import (
	"github.com/steve-care-software/webx/engine/stencils/applications/layers/instructions/assignments"
	"github.com/steve-care-software/webx/engine/stencils/applications/layers/instructions/executions"
	"github.com/steve-care-software/webx/engine/stencils/applications/layers/instructions/lists"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/interruptions/failures"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	assignmentApp assignments.Application,
	listApp lists.Application,
	execApp executions.Application,
) Application {
	stackBuilder := stacks.NewBuilder()
	framesBuilder := stacks.NewFramesBuilder()
	frameBuilder := stacks.NewFrameBuilder()
	assignmentsBuilder := stacks.NewAssignmentsBuilder()
	interruptionBuilder := interruptions.NewBuilder()
	failureBuilder := failures.NewBuilder()
	return createApplication(
		assignmentApp,
		listApp,
		execApp,
		stackBuilder,
		framesBuilder,
		frameBuilder,
		assignmentsBuilder,
		interruptionBuilder,
		failureBuilder,
	)
}

// Application represents an instructions application
type Application interface {
	Execute(frame stacks.Frame, instructions instructions.Instructions) (stacks.Frame, interruptions.Interruption, error)
}
