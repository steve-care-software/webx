package executions

import (
	"github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/executions/merges"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execMergeApp merges.Application,
) Application {
	return createApplication(
		execMergeApp,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignment executions.Execution) (*uint, error)
}
