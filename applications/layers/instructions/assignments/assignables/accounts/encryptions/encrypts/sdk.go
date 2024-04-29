package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application for tests
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, *uint, error)
}
