package validates

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		assignableBuilder,
	)
}

// Application represents a validate signature application
type Application interface {
	Execute(frame stacks.Frame, assignable validates.Validate) (stacks.Assignable, *uint, error)
}
