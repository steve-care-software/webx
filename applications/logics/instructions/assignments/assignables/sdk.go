package assignables

import (
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execCompilerApp compilers.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
	execCryptoApp cryptography.Application,
	execDatabaseApp databases.Application,
	execListApp lists.Application,
) Application {
	return createApplication(
		execCompilerApp,
		execBytesApp,
		execConstantApp,
		execCryptoApp,
		execDatabaseApp,
		execListApp,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable assignables.Assignable) (stacks.Assignable, *uint, error)
}
