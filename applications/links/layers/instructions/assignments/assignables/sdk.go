package assignables

import (
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/libraries"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execAccountApp accounts.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
	execCryptoApp cryptography.Application,
	execLibraryApp libraries.Application,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		execAccountApp,
		execBytesApp,
		execConstantApp,
		execCryptoApp,
		execLibraryApp,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable assignables.Assignable) (stacks.Assignable, *uint, error)
}
