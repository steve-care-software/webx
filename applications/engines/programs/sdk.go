package programs

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

// Application represents a program application
type Application interface {
	Execute(tree trees.Tree, command commands.Command) (programs.Program, error)
}
