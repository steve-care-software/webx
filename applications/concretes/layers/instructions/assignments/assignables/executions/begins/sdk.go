package begins

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_begins "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution begin application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_begins.Begin) (stacks.Assignable, *uint, error)
}
