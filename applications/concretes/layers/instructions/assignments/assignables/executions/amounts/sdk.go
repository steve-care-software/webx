package amounts

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_amounts "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution amount application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_amounts.Amount) (stacks.Assignable, *uint, error)
}
