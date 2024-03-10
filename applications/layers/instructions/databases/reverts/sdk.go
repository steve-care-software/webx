package reverts

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction reverts.Revert) (*uint, error)
}
