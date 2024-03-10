package accounts

import (
	instructions_accounts "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, instruction instructions_accounts.Account) (*uint, error)
}
