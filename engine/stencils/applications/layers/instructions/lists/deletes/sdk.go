package deletes

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/lists/deletes"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuiler := stacks.NewAssignableBuilder()
	assignablesBuilder := stacks.NewAssignablesBuilder()
	assignmentBuilder := stacks.NewAssignmentBuilder()
	return createApplication(
		assignableBuiler,
		assignablesBuilder,
		assignmentBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignment deletes.Delete) (stacks.Assignment, *uint, error)
}
