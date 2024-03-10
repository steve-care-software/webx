package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/commands"
	"github.com/steve-care-software/datastencil/domain/instances/libraries"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers"
)

// Application represents the logic application
type Application interface {
	Execute(
		input []byte,
		layer layers.Layer,
		library libraries.Library,
		context commands.Commands,
	) (commands.Commands, error)
}
