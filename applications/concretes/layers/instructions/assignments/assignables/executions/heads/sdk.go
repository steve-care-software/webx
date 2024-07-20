package heads

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_heads "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution head application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignable instruction_heads.Head) (stacks.Assignable, *uint, error)
}
