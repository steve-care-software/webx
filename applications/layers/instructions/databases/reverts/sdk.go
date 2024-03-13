package reverts

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/reverts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	service instances.Service,
) Application {
	return createApplication(
		service,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction reverts.Revert) (*uint, error)
}
