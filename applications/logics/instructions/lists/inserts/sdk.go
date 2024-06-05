package inserts

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
	Execute(frame stacks.Frame, assignable inserts.Insert) (stacks.Assignment, *uint, error)
}
