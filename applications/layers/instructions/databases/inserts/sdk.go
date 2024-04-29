package inserts

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	repository instances.Repository,
	service instances.Service,
) Application {
	return createApplication(
		repository,
		service,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction inserts.Insert) (*uint, error)
}
