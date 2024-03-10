package signs

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable signs.Sign) (stacks.Assignable, error)
}
