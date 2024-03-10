package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(stack stacks.Stack, instructions instructions.Instructions) (bool, stacks.Stack, results.Failure, error)
}
