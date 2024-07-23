package constants

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	assigmablesBuilder := stacks.NewAssignablesBuilder()
	return createApplication(
		assigmablesBuilder,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(constants.Constant) (stacks.Assignable, *uint, error)
}
