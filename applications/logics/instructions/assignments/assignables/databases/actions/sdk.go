package actions

import (
	database_actions "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	modificationsBuilder := modifications.NewBuilder()
	actionBuilder := database_actions.NewActionBuilder()
	return createApplication(
		assignableBuilder,
		modificationsBuilder,
		actionBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable actions.Action) (stacks.Assignable, *uint, error)
}
