package logics

import (
	"github.com/steve-care-software/datastencil/domain/commands"
	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
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
