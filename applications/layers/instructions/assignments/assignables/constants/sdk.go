package constants

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(constants.Constant) (stacks.Assignable, *uint, error)
}
