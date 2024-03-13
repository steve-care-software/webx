package deletes

import (
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/deletes"
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
	Execute(frame stacks.Frame, instruction deletes.Delete) (*uint, error)
}
