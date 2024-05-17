package databases

import (
	database_instruction "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignment database_instruction.Database) (*uint, error)
}
