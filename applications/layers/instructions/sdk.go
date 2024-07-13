package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results/interruptions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an instructions application
type Application interface {
	Execute(frame stacks.Frame, instructions instructions.Instructions) (stacks.Frame, interruptions.Interruption, error)
}
