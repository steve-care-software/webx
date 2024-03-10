package services

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable services.Service) (stacks.Assignable, error)
}
