package databases

import (
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases/retrieves"
	assignables_databases "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	actionApp actions.Application,
	commitApp commits.Application,
	databaseApp databases.Application,
	deleteApp deletes.Application,
	modificationApp modifications.Application,
	retrieveApp retrieves.Application,
) Application {
	return createApplication(
		actionApp,
		commitApp,
		databaseApp,
		deleteApp,
		modificationApp,
		retrieveApp,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable assignables_databases.Database) (stacks.Assignable, *uint, error)
}
