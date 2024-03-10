package accounts

import (
	assignables_accounts "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable assignables_accounts.Account) (stacks.Assignable, error)
}
