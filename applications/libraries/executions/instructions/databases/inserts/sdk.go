package inserts

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction inserts.Insert) (*uint, error)
}
