package accounts

import (
	"github.com/steve-care-software/datastencil/domain/commands"
	"github.com/steve-care-software/datastencil/domain/commands/results"
	"github.com/steve-care-software/datastencil/domain/libraries"
	instructions_accounts "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts"
	assignables_accounts "github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithLibrary(library libraries.Library) Builder
	WithContext(context commands.Commands) Builder
	Now() (Application, error)
}

// Application represents an execution account application
type Application interface {
	Instruction(stack stacks.Stack, instruction instructions_accounts.Account) (bool, stacks.Stack, results.Failure, commands.Commands, error)
	Assignable(frame stacks.Frame, assignable assignables_accounts.Account) (stacks.Assignable, results.Failure, commands.Commands, error)
}
