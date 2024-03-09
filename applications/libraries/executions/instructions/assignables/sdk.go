package assignables

import (
	"github.com/steve-care-software/datastencil/domain/commands"
	"github.com/steve-care-software/datastencil/domain/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables"
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
	Execute(frame stacks.Frame, assignable assignables.Assignable) (stacks.Assignable, error)
}
