package retrieves

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_retrieves "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution retrieve application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_retrieves.Retrieve) (stacks.Assignable, *uint, error)
}
