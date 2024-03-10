package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction databases.Database) (*uint, error)
}
