package inits

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_inits "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution init application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_inits.Init) (stacks.Assignable, *uint, error)
}
