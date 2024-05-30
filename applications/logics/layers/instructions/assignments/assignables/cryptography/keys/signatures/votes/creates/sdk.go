package creates

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		assignableBuilder,
	)
}

// Application represents a create vote application
type Application interface {
	Execute(frame stacks.Frame, assignable creates.Create) (stacks.Assignable, *uint, error)
}
