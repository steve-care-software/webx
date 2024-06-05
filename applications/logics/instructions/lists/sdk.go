package lists

import (
	"github.com/steve-care-software/datastencil/applications/logics/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/lists"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	insertApp inserts.Application,
	deleteApp deletes.Application,
) Application {
	return createApplication(
		insertApp,
		deleteApp,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, list lists.List) (stacks.Assignment, *uint, error)
}
