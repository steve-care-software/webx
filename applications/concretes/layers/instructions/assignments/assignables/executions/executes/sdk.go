package executes

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_executes "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		assignableBuilder,
	)
}

// Application represents an execution execute application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_executes.Execute) (stacks.Assignable, *uint, error)
}