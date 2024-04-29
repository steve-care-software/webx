package services

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/services"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	service instances.Service,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		service,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable services.Service) (stacks.Assignable, *uint, error)
}
