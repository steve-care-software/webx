package constants

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
