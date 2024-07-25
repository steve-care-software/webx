package inits

import (
	"github.com/steve-care-software/webx/engine/stencils/applications"
	instruction_inits "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		assignableBuilder,
	)
}

// Application represents an execution init application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_inits.Init) (stacks.Assignable, *uint, error)
}
