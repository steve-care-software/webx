package retrieves

import (
	"github.com/steve-care-software/datastencil/stencils/applications"
	instruction_retrieves "github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		assignableBuilder,
	)
}

// Application represents an execution retrieve application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_retrieves.Retrieve) (stacks.Assignable, *uint, error)
}
