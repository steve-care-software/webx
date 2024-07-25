package lists

import (
	application_fetches "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	fetchApplication application_fetches.Application,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	assignablesBuilder := stacks.NewAssignablesBuilder()
	return createApplication(
		fetchApplication,
		assignableBuilder,
		assignablesBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable lists.List) (stacks.Assignable, *uint, error)
}
