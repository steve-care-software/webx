package decrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
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
	Execute(frame stacks.Frame, assignable decrypts.Decrypt) (stacks.Assignable, *uint, error)
}
