package applications

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands"
)

// Application represents an application
type Application interface {
	Execute(input []byte) (commands.Commands, error)
	ExecuteWithContext(input []byte, context commands.Commands) (commands.Commands, error)
}
