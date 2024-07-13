package executions

import (
	instruction_execution "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable instruction_execution.Execution) (stacks.Assignable, *uint, error)
}
