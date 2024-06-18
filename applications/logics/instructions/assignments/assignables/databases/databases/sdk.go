package databases

import (
	databases_databases "github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	databaseBuilder := databases_databases.NewBuilder()
	headBuilder := heads.NewBuilder()
	return createApplication(
		assignableBuilder,
		databaseBuilder,
		headBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable databases.Database) (stacks.Assignable, *uint, error)
}
