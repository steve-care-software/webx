package compilers

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/instances"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	instanceAdapter instances.Adapter,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		instanceAdapter,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable compilers.Compiler) (stacks.Assignable, *uint, error)
}