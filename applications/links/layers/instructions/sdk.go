package instructions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an instructions application
type Application interface {
	Execute(stack stacks.Stack, instructions instructions.Instructions) (stacks.Stack, *uint, *uint, error)
}
