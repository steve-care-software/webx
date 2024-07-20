package merges

import (
	"github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignment merges.Merge) (*uint, error)
}
