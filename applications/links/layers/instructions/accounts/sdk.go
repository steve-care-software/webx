package accounts

import (
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/applications/links/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/domain/accounts"
	instructions_accounts "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execInsertApp inserts.Application,
	execUpdateApp updates.Application,
	service accounts.Service,
) Application {
	return createApplication(
		execInsertApp,
		execUpdateApp,
		service,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction instructions_accounts.Account) (*uint, error)
}
